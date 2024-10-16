# Product-BE GVHD thầy Lê Đình Thuận
## tech: Go, Gin, MongoDB, các thư viện khác
## Để chạy chương trình
### Cách 1
```bash
go install github.com/air-verse/air@latest
```
```bash
air
```
### Cách 2
```bash
go run main.go
```
### Cách 3
Sử dụng Docker
```bash
Docker version
```
Kiểm tra phiên bản version đang sử dụng nếu không có thì tải docker xuống
```bash
docker build -t <user>/<name>:<version> .
docker run -d -p <port_1>:<port_2> --name <name> --env-file <file .env> <user>/<name>:<version>
```