#!/bin/bash

[ "$#" -ne 3 ] && echo "parameters required GOOS GOARCH FILENAME" && exit

mkdir build >/dev/null 2>/dev/null
cp settings.conf.example build/settings.conf
env GOOS="$1" GOARCH="$2" go build -o build/"$3"