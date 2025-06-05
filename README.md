# Procwatch

A simple go binary using a for loop to monitor for file changes in `/proc`. When changes are detected, the environment variables for the new process are printed, and also written to `$PWD/proc_monitor.log`

To run the image for GitHub Actions:

```
docker run \
  -e URL=https://github.com/Your-Org-Name \
  -e TOKEN=YOUR_TOKEN \
  -e LABELS=your,comma,separated,labels \
  -e NAME=YourRunnerName
  smarticu5/procwatch:gha
```

Be aware that by default, the name of the runner is "procwatch", and the script will automatically try to replace a runner with the same name. The runners are also enrolled as ephemeral, which means they will unenroll upon completion of the first job. 
