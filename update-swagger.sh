#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
export GO111MODULE=on

# rm -f -r ../assets/swagger/docs
rm -f -r ./docs/swagger/docs

# go get -u github.com/swaggo/swag/cmd/swag
go mod download

swag init -o ./docs/swagger/docs --ot json,yaml

echo "Suceessfully generated swagger docs."
