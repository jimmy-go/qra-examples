#!/bin/sh
# run as DB_URL="your connection string" sh run.sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra/pgmanager/db/migration

migrate -path="$PKG_DIR" -url="postgres://$(echo $DB_URL)" reset
