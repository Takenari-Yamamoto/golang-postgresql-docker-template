## マイグレーション

### migration 作成

コンテナの中で

```
migrate create -ext sql -dir database/migrations -seq テーブル名
```

### migrate up

```
migrate -database postgres://${POSTGRES_USER:-postgres}:${POSTGRES_PASSWORD:-password}@db:5432/${POSTGRES_DB:-golang_postgres}?sslmode=disable -path database/migrations up
```

### migrate down

```
migrate -database postgres://${POSTGRES_USER:-postgres}:${POSTGRES_PASSWORD:-password}@db:5432/${POSTGRES_DB:-golang_postgres}?sslmode=disable -path database/migrations down
```

## テーブルからコードを自動生成

```
sqlboiler psql --output database/models --pkgname models --wipe
```

### GraphQL スキーマ生成

```
$ go run github.com/99designs/gqlgen@latest generate
```

### PostgreSQL

```
psql -p 5432 -U postgres
```

```
migrate create -ext sql -dir database/migrations -seq create_location_table
```
