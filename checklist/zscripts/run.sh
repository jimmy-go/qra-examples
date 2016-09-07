#!/bin/sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra/examples/checklist

cd $PKG_DIR/cmd/checklist

go build -o $GOBIN/qra_ex1

$GOBIN/qra_ex1 \
    -port=5050 \
    -templates=$PKG_DIR/templates \
    -static=$PKG_DIR/static \
    -db=""
