FROM cgr.dev/chainguard/go AS builder

ENV GOBIN=/usr/local/bin

RUN go install github.com/go-task/task/v3/cmd/task@latest \
 && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
 && go install github.com/pressly/goose/v3/cmd/goose@latest \
 && go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN task build

FROM cgr.dev/chainguard/glibc-dynamic

WORKDIR /app

COPY --from=builder /app/bin/go-starter ./
COPY --from=builder /app/public ./public

CMD ["./go-starter"]
