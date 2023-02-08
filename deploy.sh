#!/bin/bash
[ "$#" -ne 3 ] && echo "parameters required GOOS GOARCH FILENAME" && exit

./build.sh "$@"
mkdir deploy >/dev/null 2>/dev/null
tar -cf deploy/deploy-lesgofile-"$(date -I)"-"$1"-"$2".tar build/"$3" build/settings.conf 
