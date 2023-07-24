#!/bin/bash
SYSTEMS=("linux" "windows")
for SYSTEM in ${SYSTEMS[@]}; do
if [[ "$SYSTEM" == "windows" ]]; then
    ./deploy.sh "$SYSTEM" "amd64" lesgofile.exe
else
    ./deploy.sh "$SYSTEM" "amd64" lesgofile
fi
done