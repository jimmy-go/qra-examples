#!/bin/sh
# run as DB_URL="your connection string" sh run.sh

export PKG_DIR=$GOPATH/src/github.com/jimmy-go/qra/pgmanager/migration

migrate -path="$PKG_DIR" -url="$PGLIST_MIGRATE" reset
