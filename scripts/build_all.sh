#! /bin/bash

mkdir -p dist
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/gctu.exe
zip -j ./dist/gctu_win_amd64.zip ./dist/gctu.exe
rm ./dist/gctu.exe
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/gctu
zip -j ./dist/gctu_linux_amd64.zip ./dist/gctu
rm ./dist/gctu
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./dist/gctu
zip -j ./dist/gctu_linux_arm64.zip ./dist/gctu
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ./dist/gctu
zip -j ./dist/gctu_darwin_arm64.zip ./dist/gctu
rm ./dist/gctu
