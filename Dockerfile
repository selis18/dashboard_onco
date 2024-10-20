FROM golang:1.23.2 AS builder
    
WORKDIR /app

COPY main.go go.mod go.sum ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o server

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server ./

COPY src/ ./src/

EXPOSE 8080

CMD ["./server", "--port", "8080"]