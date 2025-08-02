package handler

import (
	"encoding/json"
	"login-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	exists, _ := Rdb.Exists(Ctx, "user:"+user.Username).Result()
	if exists == 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "ชื่อผู้ใช้นี้มีอยู่แล้ว"})
		return
	}

	userData, _ := json.Marshal(user)
	Rdb.Set(Ctx, "user:"+user.Username, userData, 0)
	c.JSON(http.StatusOK, gin.H{"message": "สมัครสมาชิกสำเร็จ"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	val, err := Rdb.Get(Ctx, "user:"+input.Username).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบผู้ใช้นี้"})
		return
	}

	json.Unmarshal([]byte(val), &user)

	if input.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้าง token ได้"})
		return
	}

	Rdb.Set(Ctx, "token:"+user.Username, tokenStr, time.Hour)

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func Logout(c *gin.Context) {
	usernameInterface, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "session หรือ token ไม่ถูกต้อง"})
		return
	}
	username, ok := usernameInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ข้อมูลผู้ใช้ในบริบทเสียหาย"})
		return
	}

	err := Rdb.Del(Ctx, "token:"+username).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดำเนินการ logout ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout สำเร็จ"})
}
