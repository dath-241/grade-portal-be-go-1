# Danh Sách Các API và Chức Năng

## Admin

- **URL**: `domain/admin/api`

### Tài Khoản Admin (Auth_Route) 
  - **Đăng Nhập**: `URL/login`
    - Phương thức: POST  
    - Mô tả: Tính năng đăng nhập, cookie sẽ được ghi vào cookie trên máy người dùng trong vòng 24h
    ```json
      request :
      {
        "idToken": string
      },
      response :
      {
        "code":  "Success",
        "token": string
      }
    ```

  - **Đăng Xuất**: `URL/logout`  
    - Phương thức: POST 
    - Mô tả: Tính năng đăng xuất, xóa cookie trên máy người dùng
    ```json
    response: 
    { 
      "code": "Success",
      "massage": "Đăng xuất thành công",
    }
    ```

  - **Tạo Tài Khoản Admin**: `URL/create`
    - Phương thức: POST 
    - Mô tả: Tính năng tạo tài khoản admin
    ```json
      request :
      {	    
        "email": string,
        "name" :  string,
        "faculty": string,
        "ms":      string,
      },
      response :
      {	    
        "code": "vao duoc trang createAdmin",
      }
    ```

  - **Thông tin tài khoản**: `URL/profile`
    - Phương thức: GET 
    - Mô tả: Xem chi tiết thông tin tài khoản
    ```json
      response :
      {	    
        "code": "success",
		    "msg": "Thanh cong",
        "user": 
        {
          "ID": string  ,
          "Email": string,
          "Name": string,
          "Ms": string,
          "Faculty": string,
          "CreatedBy": string 
        }
      }
    ```

### Quản Lý Tài Khoản Người dùng (Account_route)
  - **Tạo Tài Khoản**: `URL/account/create`  
    - Phương thức: POST
    - Mô tả: Tạo thêm tài khoản (có thể gửi lên một danh sách tài khoản).
    ```json
      request: 
      {
        [
          {
            "email": string,
            "name": string,
            "ms": string,
            "faculty": string,
            "role": string
          },
          {
            // account 2
          },
          ...
        ]
      }
      response: 
      {
        "code": "success",
		    "errorAccount":  errorAccount, //nguoi dung bi trung lap (format nhu accessAccount)
        "accessAccount": 
        [
          {
            "email": string,
            "name": string,
            "ms": string,
            "faculty": string,
            "role": string,
            "createdBy": string,
	          "expiredAt": string
          },
          //acount 2
        ]
      }
    ```

  - **Thông tin chi tiết tài khoản**: `URL/account/:id` 
    - Phương thức: GET 
    - Mô tả: Lấy chi tiết tài khoản theo accountId
    ```json
      response: 
      {
        "status": "User found successfully",
        "account": 
        {
          "ID": string,
          "Email": string,
          "Name": string,
          "Ms": string,
          "Faculty": string,
          "Role": string,
          "CreatedBy": string,
          "ExpiredAt": string
        }
      }
    ```

  - **Lấy Tài Khoản có role là Teacher**: `URL/account/teacher` 
    - Phương thức: GET 
    - Mô tả: 
      - 1.Lấy tất cả tài khoản có role là teacher
      ```json
      response: 
      {
        "status": "Users found successfully",
        "account": 
        {
         "ID": string,
          "Email": string,
          "Name": string,
          "Ms": string,
          "Faculty": string,
          "Role": string,
          "CreatedBy": string,
          "ExpiredAt": string
        }, 
        {
          //account 2
        },
        ...
      }
      ```
      - 2.Hoặc lấy 1 tài khoản có role là teacher có mã số ms bằng cách sử dụng API `URL/account/teacher?ms=?`
      ```json 
       response: 
       {
        "status": "User found successfully",
        "foundedUser": 
        {
          "ID": string,
          "Email": string,
          "Name": string,
          "Ms": string,
          "Faculty": string,
          "Role": string,
          "CreatedBy": string,
          "ExpiredAt": string
        }
      }
      ```

  - **Lấy Tài Khoản có role là Student**: `URL/account/student` 
    - Phương thức: GET 
    - Mô tả: 
      - 1.Lấy tất cả tài khoản có role là student
        ```json
        response: 
        {
          "status": "Users found successfully",
          "foundedUser": 
          {
            "ID": string,
            "Email": string,
            "Name": string,
            "Ms": string,
            "Faculty": string,
            "Role": string,
            "CreatedBy": string,
            "ExpiredAt": string
          }, 
          {
            //account 2
          },
          ...
        }
        ```
        - 2.Hoặc lấy 1 tài khoản có role là student có mã số ms bằng cách sử dụng API `URL/account/student?ms=?`
        ```json 
        response: {
          "status": "User found successfully",
          "account": 
          {
            "ID": string,
            "Email": string,
            "Name": string,
            "Ms": string,
            "Faculty": string,
            "Role": string,
            "CreatedBy": string,
            "ExpiredAt": string
          }
        } 
        ```

  - **Xoá tài khoản theo id**: `URL/account/delete/:id` 
    - Phương thức: DELETE 
    - Mô tả: Xoá tài khoản theo accountId
    ```json 
        response: 
        {
          "code": "success",
		      "msg":  "Xóa account thành công",
          "user": 
          {
            "DeletedCount": int,
            "Acknowledged": boolean
          }
        } 
      ```

  - **Chỉnh sửa tài khoản theo id**: `URL/account/change/:id` 
    - Phương thức: PATCH 
    - Mô tả: Chỉnh sửa tài khoản theo accountId
    ```json
    request: 
        {
          "name": string,
          "faculty": string,
          "role": string
        },
    response: 
        {
          "code": "success",
		      "msg": "Thay doi thanh cong"
        } 
    ```

### Quản Lý Khóa Học
  - **Tạo Khóa Học**: `URL/course/create`  
    - Phương thức: POST 
    - Mô tả: Tạo thêm 1 khóa học.
    ```json
    request: 
      {
        "ms": string,
        "name": string,
        "credit": int,
        "desc": string,
        "BT": int,
        "TN": int,
        "BTL": int,
        "GK": int,
        "CK": int
      },
    response: 
      {
        "code":    "success",
		    "message": "Tạo khóa học thành công"
      } 
    ```

  - **Lấy Khoá Học Theo Id**: `URL/course/:id`  
    - Phương thức: GET 
    - Mô tả: Lấy Khoá Học Theo Id.
    ```json
    response: 
      {
        "status":  "success",
        "message": "Lấy môn học thành công",
        "course":  
        {
          "ID": string,
          "MS": string,
          "Credit": int,
          "Name": string,
          "Desc": string,
          "HS": [5]int,
          "CreatedBy" : string,
          "Updatedby" : string
        }
      } 
    ```

  - **Lấy Tất Cả Khoá Học**: `URL/course/all`  
    - Phương thức: GET 
    - Mô tả: Lấy Tất Cả Khoá Học.
    ```json
    response: 
      {
        "code":      "success",
        "msg":       "Lấy ra tất cả khóa học thành công",
        "allCourse":  
        [
          {
            "ID": string,
            "MS": string,
            "Credit": int,
            "Name": string,
            "Desc": string,
            "HS": [5]int,
            "CreatedBy" : string,
            "Updatedby" : string
          },
          {
            // course2
          },
          ...
        ],
        "semester":  string,
      } 
    ```

  - **Xoá Khoá Học Theo Id**: `URL/delete/:id`  
    - Phương thức: Delete 
    - Mô tả: Xoá Khoá Học Theo Id.
    ```json
    response: 
      {
        "code": "success",
		    "msg":  "Xóa khóa học thành công",
      } 
    ```

  - **Chỉnh sửa Khoá Học Theo Id**: `URL/change/:id`  
    - Phương thức: PATCH 
    - Mô tả: Chỉnh sửa Khoá Học Theo Id.
    ```json
    request:
    {
      "ms": string,
      "name": string,
      "credit": int,
      "desc": string,
    },
    response: 
      {
        "code": "success",
		    "msg":  "Change course thanh cong",
      } 
    ```

### Quản Lý Lớp Học    
  - **Tạo Lớp Học**: `URL/class/create`  
    - Phương thức: POST 
    - Mô tả: Tạo thêm 1 lớp học mới.
    ```json
    request:
    {
      "name": string,
      "semester": string,
      "course_id": string,
      "teacher_id": string,
      "listStudent_id": [ // 1 mảng các string mssv
        string, string, ...
      ]
    },
    response:
    {
      "code": "success",
		  "msg":  "Tạo lớp học thành công"
    }
    ```

  - **Lấy lớp bằng id lớp ***: `URL/class/:id`
    - Phương thức: GET
    - Mô tả: Tính năng lấy ra lớp học bằng id lớp học, nếu tìm thấy lớp học trả về 
    ```json
    response:
    { 
      "status":  "success",
      "message": "Lấy lớp học thành công",
      "class":
      {
        "ID": string,           
        "Semester": string,        
        "Name": string,           
        "CourseId": string,         
        "ListStudentMs": [
          string,
          string,
          ...
        ],                
        "TeacherId": string,     
        "CreatedBy": string,     
        "UpdatedBy": string,     
      }
    }
  ```

  - **Lấy Lớp Theo ID Tài Khoản**: `URL/class/account/:id`
    - Phương thức: GET
    - Mô tả: Lấy 1 danh sách các lớp học dựa vào id (có thể là student hoặc teacher) 
     ```json
    response:
    { 
      "status":  "success",
      "message": "Lấy lớp học thành công",
      "data": 
      {
        "classes": 
        [
          {
            "_id,omitempty": string,           
            "semester": string,        
            "name": string,           
            "courseId": string,         
            "listStudentId": [
              string,
              string,
              ...
            ],                
            "teacherId": string,     
            "createdBy": string,     
            "updatedBy": string,     
          },
          {
            // class 2
          },
          ...
        ]
      }
    }
  ```

  - **Lấy Lớp Theo ID Khoá Học**: `URL/class/course/:id`
    - Phương thức: GET
    - Mô tả: Lấy 1 danh sách các lớp học dựa vào id (khoá học) 
    ```json
      response:
      { 
        "status":  "success",
        "message": "Lấy lớp học thành công",
        "data": 
        {
          "classes": 
          [
            {
              "_id,omitempty": string,           
              "semester": string,        
              "name": string,           
              "courseId": string,         
              "listStudentId": [
                string,
                string,
                ...
              ],                
              "teacherId": string,     
              "createdBy": string,     
              "updatedBy": string,     
            },
            {
              // class 2
            },
            ...
          ]
        }
      }
    ```

  - **Thêm Học Sinh Vào Lớp**: `URL/class/add`
    - Phương thức: PATCH
    - Mô tả: Thêm học sinh vào lớp
    ```json
      request:
      {
        "class_id": string,
        "listStudent_ms": 
        [
          string, 
          string,
          ...
        ]
      }
      response:
      { 
        "code": "success",
        "message": "Students added to course successfully"
      }
    ```

  - **Xoá Lớp Học Theo Id**: `URL/class/delete/:id`
    - Phương thức: DELETE
    - Mô tả: Xoá lớp học theo id
    ```json
      response:
      { 
        "code": "success",
        "msg": "xoa class thanh cong"
      }
    ```

  - **Chỉnh Sửa Lớp Học Theo Id**: `URL/class/change/:id`
    - Phương thức: PATCH
    - Mô tả: Chỉnh sửa lớp học theo id
    ```json
      request:
      {
        "semester": string,
        "name": string,
        "course_id": string,
        "teacher_id": string
      }
      response:
      { 
        "code": "success",
        "msg": "update lop hoc thanh cong",
        "class": {
          "MatchedCount": int,
          "ModifiedCount": int,
          "UpsertedCount": int,
          "UpsertedID": string,
          "Acknowledged": boolean
        },
      }
    ```

### Kết Quả Học Tập
  - **Tạo Bảng Kết Quả**: `URL/resultScore/create`  
    - Phương thức: POST
    - Mô tả: Tạo thêm 1 bảng kết quả học tập.
    ```json
    request:
      {
        "score": [
          "MSSV": string,
          "Data": {
            "BT": []float // 1 mảng các điểm BT,
            "TN": []float // 1 mảng các điểm TN,
            "BTL": []float // 1 mảng các điểm BTL,
            "GK": float // điểm giữa kỳ
            "CK": float // điểm cuối kỳ
          }
        ],
        "class_id": string,
        "course_id": string
      }
    respose:
    {
      "code": "success",
      "massage": "Cap nhat bang diem thanh cong"
    }
    ```

  - **Xem Bảng Kết Quả Theo Id**: `URL/resultScore/:id`  
    - Phương thức: GET
    - Mô tả: Xem Bảng Kết Quả Theo Id.
    ```json 
    respose:
    {
      "code": "success",
      "msg": "Lấy bảng điểm thành công",
      "score": 
      {
        "ID": string,
        "Semester": string,
        "SCORE": 
        [
          {
            "MSSV": string,
            "Data": 
            {
                "BT": []float,
                "TN": []float,
                "BTL": []float,
                "GK": float,
                "CK": float
            }
          },
          {
            //student 2
          },
          ...
        ]
      }
    }
    ```

### Hall Of Fame
  