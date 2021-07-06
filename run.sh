#!/bin/bash
go get -d -v ./... && \
go install -v ./... && \
go build && \
./clere-coding-challenge-api
