# 🎓 **Product-BE – GVHD: Thầy Lê Đình Thuận**  

## 📝 **1. Giới thiệu**  
**Product-BE** là phần mềm **backend** cho hệ thống **Grade Portal** – nền tảng đồng bộ dữ liệu bảng điểm, giúp sinh viên dễ dàng tra cứu điểm môn học.  

- **📋 Đề tài**: **Grade Portal**  
- **💻 Nhiệm vụ**: Phát triển Backend bằng **Go 1.x**  
- **👨‍💻 Nhóm**: *"Code không bao giờ khó"*  

---

## 🛠 **2. Công nghệ sử dụng**  
| 🧰 Công nghệ       | 📝 Mô tả                       |
|-------------------|--------------------------------|
| **Go (Golang)**   | Ngôn ngữ lập trình chính       |
| **Gin**           | Framework Web cho Go          |
| **MongoDB**       | Cơ sở dữ liệu NoSQL            |
| **Docker**        | Container hóa ứng dụng         |
| **GitFlow**       | Quy trình phát triển phần mềm  |
| **Git & GitHub**  | Quản lý mã nguồn và CI/CD      |

---

## ✨ **3. Tính năng chính**  
### 👥 **Các vai trò (Actors)**  
- **Admin**:  
  - Tạo tài khoản và phân quyền cho **Teacher**.  

- **Teacher**:  
  - Tạo môn học và đính kèm link CSV chứa bảng điểm.  
  - Hệ thống sẽ tự động tải về và đồng bộ dữ liệu mỗi khi bảng điểm có thay đổi.  

- **Student**:  
  - Tra cứu điểm của các môn học đã đăng ký.  

---

### ⏲ **Tính năng tự động hóa**  
- **Tự động cập nhật bảng điểm**: Hệ thống sẽ kiểm tra và tải dữ liệu mới trong vòng 6 tháng kể từ lần cập nhật gần nhất.  
- **Hall of Fame**: Lưu danh sách sinh viên có thành tích học tập xuất sắc.  

---

## 👨‍👩‍👧‍👦 **4. Thông tin nhóm**  
| STT | MSSV    | Tên thành viên      | Vai trò         |
|-----|---------|---------------------|-----------------|
| 1   | 2213104 | **Lý Vĩnh Thái**    | Product Owner (PO) |
| 2   | 2212372 | **Nguyễn Trung Nhân** | Developer      |
| 3   | 2211756 | **Lê Tuấn Kiệt**    | Developer      |
| 4   | 2213313 | **Trương Quang Thịnh** | Developer    |
| 5   | 2210348 | **Phùng Xương Cận** | Developer      |
| 6   | 2212918 | **Trương Tấn Sang** | Developer      |
| 7   | 2115036 | **Trịnh Khải Toàn** | Developer      |

---

## 📂 **5. Cấu trúc thư mục**  
```plaintext
.
├── /src         # Mã nguồn chính của dự án
├── /docs        # Tài liệu hướng dẫn và tài liệu API
│   ├── Deploy_guide.md   # Hướng dẫn triển khai
│   ├── User_guide.md     # Hướng dẫn sử dụng
│   └── API_doc.md        # Tài liệu API
├── /reports     # Báo cáo tiến độ và báo cáo cuối kỳ
│   ├── /weekly_reports   # Báo cáo tiến độ từng tuần
│   └── main_report.md    # Báo cáo môn học chính
└── /docker      # Cấu hình Docker cho dự án
```
