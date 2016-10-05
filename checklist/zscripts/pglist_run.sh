#!/bin/sh
# run as DB_URL="your connection string" sh run.sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra-examples/checklist

cd $PKG_DIR/cmd/pglist

if [ "$1" == "update" ]; then
    go get -v -u all
fi

if [ "$1" == "build" ]; then
    go build -o $GOBIN/qra_ex1
fi

echo "DB_URL=$DB_URL"

$GOBIN/qra_ex1 \
    -port=5151 \
    -templates=$PKG_DIR/templates \
    -db-url="$( echo $DB_URL )" \
    -static=$PKG_DIR/static
