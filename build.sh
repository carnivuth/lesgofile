#!/bin/bash
mkdir build >/dev/null 2>/dev/null
cp settings.conf build/settings.conf
go build -o build/lesgofile.o