
# Danh Sách Các API và Chức Năng

## Admin

### Đăng Nhập và Đăng Xuất `Thái`
- **POST** 
  - **Đăng Nhập**: `admin/api/login`  
    - Mô tả: Tính năng đăng nhập, cookie sẽ được ghi vào cookie trên máy người dùng trong vòng 24h
    - yêu cầu gửi lên :
    ```bash
      {
        "idToken": idToken
      }
    ```
  - **Đăng Xuất**: `admin/api/logout`  
    - Mô tả: Tính năng đăng xuất, xóa cookie trên máy người dùng

### Quản Lý Tài Khoản `Nhân và Thịnh`
- **POST**
  - **Tạo Admin Mới**: `admin/api/create`  
    - Mô tả: Tạo thêm 1 admin mới.
    - Yêu cầu gửi lên:
    ```bash
      {
        "email": email,
        "name": name,
        "ms": ms,
        "faculty": faculty
      }
    ```
  - **Tạo Tài Khoản**: `admin/api/account/create`  
    - Mô tả: Tạo thêm tài khoản (có thể gửi lên một danh sách tài khoản).
    - Yêu cầu gửi lên:
    ```bash
      [
        { // object 1
          "email": email,
          "name": name,
          "ms": mssv,
          "faculty": faculty,
          "role": role
        },
        { // object 2
          ....
        }
      ]
    ```

### Quản Lý Khóa Học và Lớp Học `Cận và Sang và Kiệt`
- **POST**
  - **Tạo Khóa Học**: `admin/api/course/create`  
    - Mô tả: Tạo thêm 1 khóa học.
    - Yêu cầu gửi lên: 
    ```bash
      {
        "name": name,
        "credit": credit,
        "ms": ms,
        "desc": desc
      }
    ```
    
  - **Tạo Lớp Học**: `admin/api/class/create`  
    - Mô tả: Tạo thêm 1 lớp học mới.
    - Yêu cầu gửi lên:
    ```bash
      {
        "name": name,
        "semester": semester,
        "course_id": course_id,
        "teacher_id": teacher_id,
        "listStudent_id": [ // 1 mảng các string mssv
          mssv_1, mssv_2, ...
        ]
      }
    ```

### Kết Quả Học Tập `Thái`
- **POST**
  - **Tạo Bảng Kết Quả**: `admin/api/resultScore/create`  
    - Mô tả: Tạo thêm 1 bảng kết quả học tập.
    - Yêu cầu gửi lên:
    ```bash
      {
        "score": [
          "MSSV": MSSV,
          "Data": {
            "BT": []float // 1 mảng các điểm BT,
            "TN": []float // 1 mảng các điểm TN,
            "BTL": []float // 1 mảng các điểm BTL,
            "GK": float // điểm giữa kỳ
            "CK": float // điểm cuối kỳ
          }
        ],
        "class_id": class_id
      }
    ```

### Tài Liệu
- **README.md**  
  - Người thực hiện: Toàn

- **Thái** đảm bảo logic cho cả team, thiết kế middleware đăng nhập trong 24h và validate dữ liệu từ frontend gửi về.

### Chú ý
- Các tính năng trên đã đảm bảo logic không trùng nhau giữa các trường chính. Bắt buộc phải đăng nhập mới được sử dụng các tính năng. Dữ liệu đã được validate trước khi chạy vào logic chính.

## Client

### Đăng nhập và đăng xuất và thông tin account đó
- **POST**
  - **Đăng Nhập**: `api/login`  
    - Mô tả: Tính năng đăng nhập, cookie sẽ được ghi vào cookie trên máy người dùng trong vòng 24h
    - yêu cầu gửi lên :
    ```bash
      {
        "idToken": idToken
        "role": role 
      }
    ```
  - **Đăng Xuất**: `api/logout`  
    - Mô tả: Tính năng đăng xuất, xóa cookie trên máy người dùng
  - **Thông tin**: `api/info`
    -Mô tả: Tính năng lấy ra dữ liệu của account đó
### Lấy ra các lớp học và chi tiết lớp học
- **GET**
  - **Lấy ra tất cả các lớp học cho giáo viên**: `api/class/teacher`
    - Mô tả: Tính năng lấy ra tất cả lớp học của giáo viên
    - Yêu cầu gửi lên: không cần gửi lên gì cả, chỉ cần đăng nhập bằng account có role là teacher
  - **Lấy ra tất cả các lớp học cho học sinh**: `api/class/student`
    - Mô tả: Tính năng lấy ra tất cả các lớp học của học sinh đó đang học
    - Yêu cầu gửi lên: không cần gửi lên gì cả, chỉ cần đăng nhập bằng account có role là student
  - **Lấy ra chi tiết lớp hoc**: `api/class/:id`
    - Mô tả: Tính năng lấy ra chi tiết lớp học đó
    - Yêu cầu gửi lên: đúng param nếu sai thì be sẽ không trả dữ liệu
### Lấy ra bảng điểm
- **Get**
  - **Lấy ra bảng điểm**: `api/resultScore/:class_id`
    - Mô tả: Tính năng lấy ra bảng điểm của lớp học đó, nếu teacher thì sẽ gửi về toàn bộ bản điểm, nếu student thì gửi về bảng điểm của student đó
    - Yêu cầu gửi lên: đúng param
  - **Lấy lớp bằng id lớp 'Nhân'**: `api/class/:id`
    - Mô tả: Tính năng lấy ra lớp học bằng id lớp học, nếu tìm thấy lớp học trả về 
    ```bash
    c.JSON(200, gin.H
   { "status":  "success",
		"message": "Lấy lớp học thành công",
		"class":{
              "ID": ""           
	            "Semester": ""        
	            "Name": ""             
	            "CourseId": ""         
	            "ListStudentId": [
                ""
                ""
                ...
                ""
              ]                   []    
	            "TeacherId": ""     
	            "CreatedBy": ""     
	            "UpdatedBy": ""     
            }
    })
    - Yêu cầu gửi lên: đúng param