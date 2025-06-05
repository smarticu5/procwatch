#!/bin/bash
set -e
act_runner register --instance $URL --token $TOKEN --name $NAME --labels $LABELS --no-interactive 

/procwatch &

act_runner daemon
