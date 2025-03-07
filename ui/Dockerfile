# Stage 1: Build React frontend
FROM node:18-alpine AS frontend-builder
WORKDIR /app
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app
# 필요한 패키지 설치
RUN apk add --no-cache git
# 소스 코드 복사
COPY . .
COPY --from=frontend-builder /app/build ./frontend/build
# Go 모듈 초기화 및 빌드
RUN go mod init ui
RUN go get github.com/gin-gonic/gin@v1.9.1
RUN go build -o ui

# Stage 3: Final image
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/ui .
COPY --from=frontend-builder /app/build ./frontend/build
EXPOSE 8080
CMD ["./ui"]
