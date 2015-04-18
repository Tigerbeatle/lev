#!/usr/bin/env bash
export MGO_HOSTS=localhost
export MGO_DATABASE=le
export MGO_USERNAME=tigerbeatle
export MGO_PASSWORD=welcome
export BUOY_DATABASE=tigerbeatle

cd $GOPATH/src/github.com/tigerbeatle/le
go clean -i
go build

bee run watchall