### Clone dự án về máy sử dụng git
```bash
git clone https://github.com/dath-241/grade-portal-be-go-1.git
```
```bash
cd /src
```

### Để chạy chương trình

#### Cách 1
```bash
go install github.com/air-verse/air@latest
```
```bash
air
```
#### Cách 2
```bash
go run main.go
```
#### Cách 3
Sử dụng Docker
```bash
Docker version
```
Kiểm tra phiên bản version đang sử dụng nếu không có thì tải docker xuống
```bash
docker build -t <user>/<name>:<version> .
docker run -d -p <port_1>:<port_2> --name <name> --env-file <file .env> <user>/<name>:<version>
```