FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init result-mapper && go build -o mapper

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mapper .
EXPOSE 8080
CMD ["./mapper"]