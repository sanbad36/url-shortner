# Build stage
FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /url-shortener

# Run stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /url-shortener /url-shortener

EXPOSE 8080

CMD ["/url-shortener"]
