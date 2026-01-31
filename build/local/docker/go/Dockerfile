FROM golang:1.25

WORKDIR /go/app

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest