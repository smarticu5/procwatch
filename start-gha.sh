#!/bin/bash
set -e
/runnertmp/config.sh --url $URL --token $TOKEN --name $NAME --labels $LABELS --unattended --replace --ephemeral --no-default-labels

/procwatch &

/home/runner/run.sh
