package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func SetJWTSecret(secret []byte) {
	jwtSecret = secret
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(jwtSecret) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET ไม่ได้ตั้งค่า"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่ได้ส่ง token มา"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "รูปแบบ Authorization ไม่ถูกต้อง"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token ไม่ถูกต้อง หรือหมดอายุ"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "claims ไม่ถูกต้อง"})
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบข้อมูลผู้ใช้ใน token"})
			c.Abort()
			return
		}

		storedToken, err := Rdb.Get(Ctx, "token:"+username).Result()
		if err != nil || storedToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token หมดอายุ หรือถูก logout"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
