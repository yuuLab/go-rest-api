# go-rest-api

## Overview

Go (echo) の Web API スターターキット。
ライブラリ実験等で使用する。

## db migrate

sample command

```
export MYSQL_URL='YOUR_URL'
export MIGRATION_PATH='/go/app/internal/infrastructure/db/migrations'

migrate create -ext sql -dir $MIGRATION_PATH -seq initialize_schema

migrate -database $MYSQL_URL -path $MIGRATION_PATH up
migrate -database $MYSQL_URL -path $MIGRATION_PATH down
```
