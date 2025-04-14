# Gunakan image Go terbaru
FROM golang:latest
# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPATH=/go \
    GOBIN=/go/bin \
    PATH=/go/bin:/usr/local/go/bin:$PATH

# Install Air
RUN go install github.com/cosmtrek/air@v1.43.0

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu
COPY go.mod go.sum ./

# Download dependency
RUN go mod download

# Copy semua file proyek ke dalam container
COPY . .

# Expose port aplikasi
EXPOSE 8080

# Jalankan aplikasi dengan Air
CMD ["air"]