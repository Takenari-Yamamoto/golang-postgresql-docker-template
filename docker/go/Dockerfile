FROM golang:1.18-alpine

RUN apk --no-cache add bash tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk add --no-cache gcc musl-dev && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 && \
    go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
