#!/usr/bin/env bash

go-bindata data/
go run nimboos.go bindata.go "$@"