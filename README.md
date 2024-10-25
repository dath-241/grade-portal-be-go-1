## Product-BE GVHD thầy Lê Đình Thuận

## Giới thiệu:
Phần mềm backend cho chương trình web để đồng bộ dữ liệu bảng điểm - cho SV có thể tra cứu điểm môn học trong quá trình giảng dạy. 

- Đề tài: Grade Portal
- Nhiệm vụ: Backend Go 1
- Tên nhóm: `Code không bao giờ khó`

### Tech: Go, Gin, MongoDB, các thư viện khác

### Tính năng:
- Actors: Admin/Teacher/Student
- Teacher có thể tạo môn học, và đính kèm là link csv chứa bảng điểm. Hệ thống sẽ tự động monitor (download nếu có phiên bản mới) và cập nhật điểm vào CSDL tra cứu.
- Admin có thể tạo/phân quyền cho Teacher
- Student có thể tra cứu điểm
- Việc cập nhật bảng điểm tự động có thời hạn là 6 tháng,
- Hall of Fame

### Thông tin nhóm
|STT | MSSV    | Tên thành viên      | Role | 
|----|---------|---------------------|------|
|1   | 2213104 | Lý Vĩnh Thái        | PO   |
|2   | 2212372 | Nguyễn Trung Nhân   | Dev  |
|3   | 2211756 | Lê Tuấn Kiệt        | Dev  |
|4   | 2213313 | Trương Quang Thịnh  | Dev  |
|5   | 2210348 | Phùng Xương Cận     | Dev  |
|6   | 2212918 | Trương Tấn Sang     | Dev  |
|7   | 2115036 | Trịnh Khải Toàn     | Dev  |

### Cấu trúc thư mục
- `/src`: Source code của dự án.
- `/docs`: Thư mục chứa các tài liệu hướng dẫn dự án, bao gồm:
  - `Deploy_guide.md`: Hướng dẫn triển khai hệ thống.
  - `User_guide.md`: Hướng dẫn sử dụng phần mềm.
  - `API_doc.md`: Tài liệu về API được sử dụng trong dự án.
- `/reports`: Thư mục chứa các báo cáo về tiến độ và báo cáo cuối kỳ.
  - `/weekly reports`: Thư mục chứ các báo cáo tiến độ các tuần
  - `main_report.md`: Báo cáo môn học.
