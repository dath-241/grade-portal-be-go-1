# Danh Sách Các API và Chức Năng

## Admin

### Đăng Nhập và Đăng Xuất `Thái`
- **POST** 
  - **Đăng Nhập**: `admin/api/login`  
    - Mô tả: Tính năng đăng nhập, yêu cầu gửi lên `idToken`.
  - **Đăng Xuất**: `admin/api/logout`  
    - Mô tả: Tính năng đăng xuất.

### Quản Lý Tài Khoản
- **POST**
  - **Tạo Admin Mới**: `admin/api/create`  
    - Mô tả: Tạo thêm 1 admin mới.
  - **Tạo Tài Khoản**: `admin/api/account/create`  
    - Mô tả: Tạo thêm tài khoản (có thể gửi lên một danh sách tài khoản).

### Quản Lý Khóa Học và Lớp Học
- **POST**
  - **Tạo Khóa Học**: `admin/api/course/create`  
    - Mô tả: Tạo thêm 1 khóa học.
  - **Tạo Lớp Học**: `admin/api/class/create`  
    - Mô tả: Tạo thêm 1 lớp học mới.

### Kết Quả Học Tập
- **POST**
  - **Tạo Bảng Kết Quả**: `admin/api/resultScore/create`  
    - Mô tả: Tạo thêm 1 bảng kết quả học tập.

### Tài Liệu
- **README.md**  
  - Người thực hiện: Toàn

## Chú Ý
- **Thái** đảm bảo logic cho cả team, thiết kế middleware đăng nhập trong 24h và validate dữ liệu từ frontend gửi về.
- Các tính năng trên đã đảm bảo logic không trùng nhau giữa các trường chính. Bắt buộc phải đăng nhập mới được sử dụng các tính năng. Dữ liệu đã được validate trước khi chạy vào logic chính.