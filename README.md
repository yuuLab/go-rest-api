# go-rest-api

Go (echo) の Web API スターターキット。
ライブラリ実験等で使用する。

# db migrate

sample command

```
export MYSQL_URL='YOUR_URL'
export MIGRATION_PATH='/go/app/internal/infrastructure/db/migrations'

migrate create -ext sql -dir $MIGRATION_PATH -seq initialize_schema

migrate -database $MYSQL_URL -path $MIGRATION_PATH up
migrate -database $MYSQL_URL -path $MIGRATION_PATH down
```

export MYSQL_URL='mysql://app_user:app_pass@tcp(rdb:3306)/app_db'
export MIGRATION_PATH='/go/app/internal/infrastructure/db/migrations'
