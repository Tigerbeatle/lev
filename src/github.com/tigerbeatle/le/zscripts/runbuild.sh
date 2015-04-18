#!/usr/bin/env bash
export MGO_HOSTS=localhost:27017
export MGO_DATABASE=le
export MGO_USERNAME=
export MGO_PASSWORD=

cd $GOPATH/src/github.com/tigerbeatle/le
go clean -i
go build

./le