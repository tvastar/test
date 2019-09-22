#!/bin/bash

export GO111MODULE=on
go run github.com/golangci/golangci-lint/cmd/golangci-lint run -E goimports -E gosec -E interfacer -E maligned -E misspell -E nakedret -E unconvert $*
