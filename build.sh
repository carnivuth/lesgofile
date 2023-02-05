#!/bin/bash
mkdir build >/dev/null 2>/dev/null
cp settings.conf build/setting.conf
go build -o build/lesGoFile