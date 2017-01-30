#!/bin/sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra-examples/checklist

go build -o $GOBIN/qra_ex1 cmd/pglist/main.go

$GOBIN/qra_ex1 \
    -port=5151 \
    -templates=$PKG_DIR/templates \
    -db-url="$PGLIST_URL" \
    -static=$PKG_DIR/static
