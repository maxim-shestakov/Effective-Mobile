#!/bin/bash

set -e 

echo $GOOSE_DBSTRING
goose -dir db/migrations up

/main
