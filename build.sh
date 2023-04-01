#!/bin/bash

[ "$#" -ne 3 ] && echo "parameters required GOOS GOARCH FILENAME" && exit

mkdir build >/dev/null 2>/dev/null
cp settings.example.json build/settings.json
env GOOS="$1" GOARCH="$2" go build -o build/"$3"