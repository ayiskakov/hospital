# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
RUN apk add build-base
COPY . .
RUN go build -o main ./cmd/srv/main.go
RUN ls

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY ./migrations migrations
COPY ./docker/config.yml config.yml
EXPOSE 8000
CMD [ "./main" ]
