FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o myapp .

FROM alpine:latest

RUN apk add --no-cache

COPY --from=builder /app/myapp /app/myapp

EXPOSE ${API_REST_PORT}

ENTRYPOINT ["/app/myapp"]