package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func readEnviron(pid string) map[string]string {
	environPath := filepath.Join("/proc", pid, "environ")
	
	data, err := os.ReadFile(environPath)
	if err != nil {
		return nil
	}

	environ := make(map[string]string)
	for _, envVar := range strings.Split(string(data), "\x00") {
		if envVar == "" {
			continue
		}
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) == 2 {
			environ[parts[0]] = parts[1]
		}
	}

	return environ
}

func formatEnviron(environ map[string]string) string {
	if len(environ) == 0 {
		return "No environment variables found"
	}

	var builder strings.Builder
	for key, value := range environ {
		builder.WriteString(fmt.Sprintf("  %s = %s\n", key, value))
	}
	return builder.String()
}

func main() {
	// Create log file
	logFile, err := os.OpenFile("proc_monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Create a multi-writer to log to both console and file
	logger := log.New(logFile, "", log.LstdFlags)
	
	procPath := "/proc"
	knownDirs := make(map[string]bool)

	initialDirs, err := os.ReadDir(procPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range initialDirs {
		if dir.IsDir() {
			knownDirs[dir.Name()] = true
		}
	}

	for {
		currentDirs, err := os.ReadDir(procPath)
		if err != nil {
			log.Printf("Error reading /proc: %v", err)
			continue
		}

		for _, dir := range currentDirs {
			if dir.IsDir() && !knownDirs[dir.Name()] {
				dirPath := filepath.Join(procPath, dir.Name())
				cmdPath := filepath.Join(dirPath, "comm")
				
				cmdContent, err := os.ReadFile(cmdPath)
				processName := "Unknown"
				if err == nil {
					processName = strings.TrimSpace(string(cmdContent))
				}

				// Log to console and file
				outputMessage := fmt.Sprintf("New Process Found:\n  PID: %s\n  Process Name: %s\n  Environment Variables:\n%s", 
					dir.Name(), 
					processName, 
					formatEnviron(readEnviron(dir.Name())))

				fmt.Println(outputMessage)
				logger.Println(outputMessage)

				knownDirs[dir.Name()] = true
			}
		}

		time.Sleep(5 * time.Millisecond)
	}
}
