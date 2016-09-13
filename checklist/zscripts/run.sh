#!/bin/sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra-examples/checklist

cd $PKG_DIR/cmd/checklist

if [ "$1" == "update" ]; then
    go get -v -u all
fi

if [ "$1" == "build" ]; then
    go build -o $GOBIN/qra_ex1
fi


$GOBIN/qra_ex1 \
    -port=5151 \
    -templates=$PKG_DIR/templates \
    -static=$PKG_DIR/static
