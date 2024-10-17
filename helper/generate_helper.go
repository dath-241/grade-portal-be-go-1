package helper

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email                string `json:"email"` // Email được lưu trong token
	jwt.RegisteredClaims        // Thêm các trường chuẩn như exp, iat
}

<<<<<<< HEAD
// tao JWT
=======
>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
func CreateJWT(email string) string {
	claims := Claims{
		Email: email, // Sử dụng email ở đây
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Hết hạn sau 1 ngày
		}, // Thời gian hết hạn của token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret) // Tạo token
	if err != nil {
		fmt.Println("Error creating token:", err)
		return ""
	}
	return tokenString
}

<<<<<<< HEAD
// dich JWT tu token
=======
>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
func ParseJWT(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET") // Lấy secret từ .env

	// Parse token với claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("phương thức ký không hợp lệ")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	// Kiểm tra nếu token hợp lệ và ép kiểu claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token không hợp lệ")
	}
}
