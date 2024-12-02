### Kết nối Database
- Chạy MongoDB Compass -> Connection string nằm bên trong file `.env`

### Test Feature bằng Postman Agent
- Import file `Backend Golang.postman_collection.json` 
- Toàn bộ API + Function + Data cần thiết đều đã nằm bên trong.
  - Đối với Login thì cần gửi vào body idToken đã lấy được sau khi đăng nhập bằng oauth ở trang index.html
![postman-agent]
  - Tất cả những hàm còn lại đều cần Authorization (`Bearer token`) (được tạo ra từ JWT_SECRET + idToken) để phân role cho người dùng