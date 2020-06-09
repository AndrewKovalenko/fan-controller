#!/bin/bash
env GOOS=linux GOARCH=arm GOARM=5 go build -o build/fan-controller ./src/cmd/fan-controller-terminal-app.go

cp ./src/cmd/fan-controller-config.yaml build/
cp ./src/cmd/start.sh build/