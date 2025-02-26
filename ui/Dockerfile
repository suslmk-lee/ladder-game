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
COPY . .
COPY --from=frontend-builder /app/build ./frontend/build
RUN go mod download
RUN go build -o ui

# Stage 3: Final image
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/ui .
COPY --from=frontend-builder /app/build ./frontend/build
EXPOSE 8080
CMD ["./ui"]
