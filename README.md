## マイグレーション

### migration 作成

```
migrate create -ext sql -dir database/migrations -seq create_テーブル名_table
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

### スキーマ生成

```
$ go run github.com/99designs/gqlgen@latest generate
```
