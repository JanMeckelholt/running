#!/bin/bash
FILE=common/grpc/token/token.pb.go
g () {
  sed "s/json:\"$1,omitempty\"/json:\"$1,omitempty\" gorm:\"$2\"/"
}

cat $FILE \
| g "clientId" "primaryKey" \
> $FILE.tmp && mv $FILE{.tmp,}