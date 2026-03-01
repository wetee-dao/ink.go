#!/bin/sh

SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

mkdir -p "$DIR/pallets"
cd "$DIR/pallets/"

curl -L -X POST -H "Content-Type: application/json" -d '{"id":1, "jsonrpc":"2.0", "method": "state_getMetadata"}' http://127.0.0.1:9944 > local-meta.json

go-substrate-gen local-meta.json "github.com/wetee-dao/ink.go/pallet"