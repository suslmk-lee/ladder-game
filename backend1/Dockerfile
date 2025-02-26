FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init ladder-manager && go build -o manager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/manager .
EXPOSE 8080
CMD ["./manager"]