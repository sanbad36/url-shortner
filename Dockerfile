# Build stage
FROM golang:1.17-alpine AS builder

WORKDIR /app

# Enable Go modules and set the Go proxy to avoid fetching issues
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

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
