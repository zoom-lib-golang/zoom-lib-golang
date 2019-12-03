#!/bin/bash

go get -u -v ./...
go mod tidy
go mod verify
go mod vendor
