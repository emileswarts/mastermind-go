#!/bin/bash

go get -u
go mod tidy
go build main.go 
go run main.go 