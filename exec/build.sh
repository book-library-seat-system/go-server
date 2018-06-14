#!/bin/bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o exec/windowsexec.exe main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exec/linuxexec main.go
go build -o exec/macexec main.go
