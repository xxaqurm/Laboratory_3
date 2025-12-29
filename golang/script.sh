#!/bin/bash

go test ./... -coverprofile=coverage.txt
go tool cover -html coverage.txt -o index.html

cp -rf cpp/coverage_html ./coverage_cpp
cp -f go/index.html ./coverage_go.html