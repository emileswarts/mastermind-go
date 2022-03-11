#!/bin/bash

go get -u
go mod tidy
go build main.go presenter.go scorer.go tick.go
go run main.go presenter.go scorer.go tick.go