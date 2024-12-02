# Danh Sách Các API và Chức Năng

## Client

-**URL** `domain/api`

### Tài khoản (Account_Route)

- **Đăng nhập web**: `URL/login`
  - Phương thức: POST
  - Mô tả: Tính năng đăng nhập cho web, cookie sữ được ghi vòa cookie trên máy người dùng trong vòng 24h.

```json
     request :{
       "idToken": string
      }
     response :{
       "code":  "Success",
		    "token": string,
		    "role": string
      }
```

- **Đăng nhập Telegram**: `URL/logintele`

  - Phương thức: POST
  - Mô tả: Tính năng đăng nhập cho Telegram, cookie sữ được ghi vòa cookie trên máy người dùng trong vòng 24h.

  ```json
      request :{
        "ms": string,
        "password": string
      }
      response :{
        "code":       "success",
		    "token":      string,
		    "listCourse": []string,
		    "role":       string
     }
  ```

- **Đăng xuất**: `URL/logout`

  - Phương thức: POST
  - Mô tả: Tính năng đăng xuất, xóa cookie trên máy người dùng.

  ```json
      response :{
        "code":    "Success",
		    "massage": "Đăng xuất thành công",
      }
  ```

- **Lấy thông tin tài khoản**: `URL/info`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra dữ liệu của account đó.
  - Giá trị trả về:
    ```json
    response :{
      "code":    "success",
      "user": {
          "_id,omitempty": bson.ObjectID,
          "email" : string ,
          "name" : string,
          "ms" : string,
          "faculty" : string,
          "role" : string,
          "createdBy" : bson.ObjectID,
          "expiredAt" : time.Time
      }
    }
    ```

- **Kiểm tra chức vụ của người dùng**: `URL/:id`

  - Phương thức: GET
  - Mô tả: Tính năng kiểm tra thử chức vụ của người dùng có phải giáo viên hay không, nếu đúng trả về thông tin giáo viên đó.
  - Yêu cầu gửi lên: param đúng id tài khoản.
  - Giá trị trả về:

  ```json
    response :{
      "code":    "success",
  	  "teacher": {
          "_id,omitempty": bson.ObjectID,
          "email" : string ,
          "name" : string,
          "ms" : string,
          "faculty" : string,
          "role" : string,
          "createdBy" : bson.ObjectID,
          "expiredAt" : time.Time
      }
    }
  ```

- **Gửi OTP**: `URL/otp`

  - Phương thức: POST
  - Mô tả: Tính năng gửi otp có giá trị 5 phút về email cho người dùng, để xác thực khi thay đổi mật khẩu.

  ```json
      request :{
        "ms": string,
      }
      response :{
       "code": "success",
  	   "msg":  "Gửi OTP thành công",
     }
  ```

- **Đổi mật khẩu**: `URL/resetpassword`

  - Phương thức: POST
  - Mô tả: Tính năng dùng để thay đổi mật khẩu của người dùng.

  ```json
      request :{
        "ms": string,
        "password": string,
        "otp":string
      }
      response :{
        "code": "success",
  	    "msg":  "Thay đổi password thành công",
     }
  ```

### Lớp học (class_route)

- **Lấy thông tin lớp học của người dùng**: `URL/class/account`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra dữ liệu tất cả các lớp của người dùng theo chức vụ.
  - Giá trị trả về:

    - Nếu là sinh viên:

    ```json
        response :{
          "code": "success",
          "classAll":[
            {
              "_id,omitempty": bson.ObjectID,
              "semester": string,
              "name": string,
              "course_id": bson.ObjectID,
              "teacher_id": bson.ObjectID
            },
            {

            }
         ]
      }
    ```

    -Nếu là giáo viên:

    ```json
       response :{
         "code": "success",
         "classAll":[
           {
             "_id,omitempty": bson.ObjectID,
             "semester": string,
             "name": string,
             "course_id": bson.ObjectID,
             "listStudent_ms": []string,
             "teacher_id": bson.ObjectID,
             "createdBy": bson.ObjectID,
             "updatedBy": bson.ObjectID
           },
           {

           }
         ]
     }
    ```

- **Lấy thông tin lớp học của người dùng**: `URL/class/:id`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra dữ liệu chi tiết của lớp học theo id lớp học.
  - Yêu cầu gửi lên: param đúng id lớp học.
  - Giá trị trả về:

    ```json
       response :{
         "code": "success",
         "classDetail" :{
             "_id,omitempty": bson.ObjectID,
             "semester": string,
             "name": string,
             "course_id": bson.ObjectID,
             "listStudent_ms": []string,
             "teacher_id": bson.ObjectID,
             "createdBy": bson.ObjectID,
             "updatedBy": bson.ObjectID
         },
     }
    ```

- **Lấy thông tin lớp học của người dùng**: `URL/class/count/:id`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra số lớp học của môn học theo id môn học.
  - Yêu cầu gửi lên: param đúng id môn học.
  - Giá trị trả về:

    ```json
       response :{
        "code": "success",
        "count": int
     }
    ```

### Khóa học (course_route)

- **Lấy thông tin lớp học của người dùng**: `URL/course/:id`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra dữ liệu chi tiết của môn học theo id môn học.
  - Yêu cầu gửi lên: param đúng id môn học.
  - Giá trị trả về:

    ```json
       response :{
        "status":  "success",
        "message": "Lấy môn học thành công",
        "course":{
          "_id,omitempty": bson.ObjectID,
          "ms": string,
          "credit": string,
          "name": string,
          "desc": string,
          "hs": [5]int,
          "createdby": bson.ObjectID,
          "updatedby": bson.ObjectID
        }
     }
    ```

### Vinh danh (hallOfFame_route)

- **Lấy thông tin lớp học của người dùng**: `URL/HOF/all`

  - Phương thức: GET
  - Mô tả: Tính năng lấy ra dữ liệu chi tiết của môn học theo id môn học.
  - Yêu cầu gửi lên: param đúng id môn học.
  - Giá trị trả về:

    ```json
       response :{
        "status":  "success",
        "message": "Lấy hall of fame thành công",
        "data":{
          "semester": string,
          "tier":[
            {
              "course_id" : bson.ObjectID,
              "Data":[
                {
                "mssv" : string ,
                "dtb" : float32
                },
                {

                }
              ]
            },
            {

            }
          ]
        }
     }
    ```

### Bảng điểm (result_route)

- **Tạo bảng điểm chỉ dành cho giáo viên**: `URL/resultScore/create`

  - Phương thức: POST
  - Mô tả: Tính năng dùng để tạo bảng điểm.

  ```json
     request :{
       "semester":  string,
  	    "course_id": bson.ObjectID,
  	    "score":[
         {
           "mssv": string,
           "data": {
             "BT": []float32,
             "TN": []float32,
             "BTL": []float32,
             "GK": float32,
             "CK": float32
           }
         },
         {

         }
       ],
  	    "class_id":  bson.ObjectID,
  	    "expiredAt": time.time,
  	    "createdBy": bson.ObjectID,
  	    "updatedBy": bson.ObjectID
     }
     response :{
       "code":    "success",
  	    "massage": "Cap nhat bang diem thanh cong"
    }
  ```

- **Cập nhật, chỉnh sửa bảng điểm chỉ dành cho giáo viên**: `URL/resultScore/:id`

  - Phương thức: PATCH
  - Mô tả: Tính năng dùng để cập nhật, chỉnh sửa bảng điểm.
  - Yêu cầu gửi lên: param đúng id bảng điểm.

  ```json
     request :{
  	    "score":[
         {
           "mssv": string,
           "data": {
             "BT": []float32,
             "TN": []float32,
             "BTL": []float32,
             "GK": float32,
             "CK": float32
           }
         },
         {

         }
       ],
  	    "updatedBy": bson.ObjectID,
     }
     response :{
       "code":    "success",
  	    "massage": "Thay đổi thành công",
    }
  ```

- **Lấy dữ liệu bảng điểm**: `URL/resultScore/:id`

  - Phương thức: GET
  - Mô tả: Tính năng dùng để lấy dữ liệu bảng điểm.
  - Yêu cầu gửi lên: param đúng id bảng điểm.
  - Giá trị trả về:

  - Nếu là sinh viên:

  ```json
      response :{
        "code": "success",
         "score":
        {
          "mssv": string,
          "data": {
            "BT": []float32,
            "TN": []float32,
            "BTL": []float32,
            "GK": float32,
            "CK": float32
          }
        }
    }
  ```

  -Nếu là giáo viên:

  ```json
     response :{
       "code": "success",
       "resultScore":{
          "_id,omitempty": bson.ObjectID,,
          "semester":  string,
  	      "course_id": bson.ObjectID,
  	      "score":[
            {
            "mssv": string,
            "data": {
              "BT": []float32,
              "TN": []float32,
              "BTL": []float32,
              "GK": float32,
              "CK": float32
              }
            },
            {

            }
          ],
  	      "class_id":  bson.ObjectID,
  	      "expiredAt": time.time,
  	      "createdBy": bson.ObjectID,
  	      "updatedBy": bson.ObjectID
        }
     }
  ```

- **Lấy dữ liệu bảng điểm của tất cả môn học**: `URL/resultScore/getmark`

  - Phương thức: GET
  - Mô tả: Tính năng dùng để lấy dữ liệu bảng điểm của tất cả môn học.

  ```json
     response :{
        "code":   "success",
  	    "msg":    "Lấy điểm thành công",
        "scores":[
          {
            "ms":string,
            "name": string,
            "data": {
                "BT": []float32,
                "TN": []float32,
                "BTL": []float32,
                "GK": float32,
                "CK": float32
              }
          },
          {

          }
        ]
     }
  ```

  - **Lấy dữ liệu bảng điểm của từng môn học**: `URL/resultScore/getmark/:ms`
  - Phương thức: GET
  - Mô tả: Tính năng dùng để lấy dữ liệu bảng điểm của từng môn học theo mã số.
  - Yêu cầu gửi lên: param đúng mã số môn học.

  ```json
     response :{
        "code":  "success",
  			"msg":   "Lấy điểm thành công",
  			"name":  string,
        "score":
        {
          "BT": []float32,
          "TN": []float32,
          "BTL": []float32,
          "GK": float32,
          "CK": float32
        }
     }
  ```
