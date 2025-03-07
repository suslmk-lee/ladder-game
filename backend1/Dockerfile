FROM golang:1.21-alpine AS builder
WORKDIR /app
# 필요한 패키지 설치
RUN apk add --no-cache git
# 소스 코드 복사
COPY . .
# Go 모듈 초기화 및 빌드
RUN go mod init ladder-manager
RUN go get github.com/gin-gonic/gin@v1.9.1
RUN go build -o manager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/manager .
EXPOSE 8080
CMD ["./manager"]