FROM golang:1.24.1 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
RUN chmod +x /app/main  # Устанавливаем права на выполнение

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
