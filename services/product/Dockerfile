# Build Stage
FROM golang:1.19.5-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go


# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

COPY pkg/config/envs/dev.env pkg/config/envs/dev.env
COPY wait-for.sh .

CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]