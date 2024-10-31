# Bước 1: Sử dụng image Go chính thức
FROM golang:1.23.0 AS builder

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum trước để tối ưu hóa cache
COPY go.* ./

# Tải các dependencies
RUN go mod download

# Sao chép mã nguồn vào thư mục làm việc
COPY . .

# Biên dịch ứng dụng với cờ để tắt cgo và chỉ định OS là linux
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-s -w' -o main .

# Bước 2: Tạo image nhẹ hơn để chạy ứng dụng
FROM alpine:latest

# Thiết lập thư mục làm việc
WORKDIR /root/

# Sao chép binary và chứng chỉ vào image mới
COPY --from=builder /app/main .


# Mở cổng mà ứng dụng sẽ lắng nghe
EXPOSE 8080

# Chạy ứng dụng khi container khởi động
CMD ["./main"]
