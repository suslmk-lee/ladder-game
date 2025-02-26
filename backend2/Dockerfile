FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init ladder-generator && go build -o generator

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/generator .
EXPOSE 8080
CMD ["./generator"]