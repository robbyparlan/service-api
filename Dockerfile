# Gunakan image base Golang versi terbaru
FROM golang:1.23-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum untuk mengunduh dependencies
COPY go.mod go.sum ./

# Unduh dependencies
RUN go mod download

# Copy seluruh kode aplikasi ke dalam container
COPY . .

# Build aplikasi Golang
RUN CGO_ENABLED=0 GOOS=linux go build -o /svc-api

# Gunakan image minimal untuk production
FROM alpine:latest

# Copy binary hasil build dari stage builder
COPY --from=builder /svc-api /svc-api

# Copy file konfigurasi
COPY config.yml /app/config.yml

# Set working directory di dalam container
WORKDIR /app

# Expose port yang akan digunakan oleh aplikasi
EXPOSE 5004

# Jalankan aplikasi saat container dijalankan
CMD ["/svc-api"]