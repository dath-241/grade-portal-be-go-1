# Danh sách các API và chức năng
## Admin  
- **POST**  
 -- **Nhân** "admin/api/login": Tính năng đăng nhập, yêu cầu gửi lên idToken  
 -- **Thái** "admin/api/logut": Tính năng đăng xuất  
 -- **Thái** "admin/api/create": Tạo thêm 1 admin mới  
 -- **Thịnh** "admin/api/account/create": Tạo thêm account (có thể gửi lên 1 danh sách account)  
 -- **Kiệt** "admin/api/course/create": Tạo thêm 1 khoá học  
 -- **Nhân** "admin/api/class/create": Tạo thêm 1 lớp học mới  
 -- **Thái** "admin/api/resultScore/creat": Tạo thêm 1 bảng kết quả học tập  
 -- **Toàn** viết file README.md  
**Thái** đảm bảo logic cho cả team, thiết kế middlewares đăng nhập trong 24h, validate dữ liệu từ frontend gửi về  
Các tính năng trên đã đảm bảo logic không trùm nhau giữa các trường chính, bất buộc phải đăng nhập mới được sử dụng các tính năng, dữ liệu đã được validate trước khi chạy vào logic chính

- 