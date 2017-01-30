#!/bin/sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra-examples/checklist

go build -o $GOBIN/qra_ex1_term cmd/pgterm/main.go

$GOBIN/qra_ex1_term \
    -db-url="$PGLIST_URL"
