#!/bin/bash
SYSTEMS=("linux" "windows")
for SYSTEM in ${SYSTEMS[@]}; do
    ./deploy.sh "$SYSTEM" "amd64" lesgofile
done