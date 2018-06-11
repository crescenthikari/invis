#!/bin/bash
echo "Start build macOS 64 bit"
GOOS=darwin GOARCH=amd64 go build -o build/macos/64/invis
echo  "Completed"
echo "Start build linux 32 bit"
GOOS=linux GOARCH=386 go build -o build/linux/32/invis
echo "Completed"
echo "Start build linux 64 bit"
GOOS=linux GOARCH=amd64 go build -o build/linux/64/invis
echo "Completed"
