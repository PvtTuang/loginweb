# 🔐 Login Backend with Go (Gin) + JWT + Redis

ระบบ **สมัครสมาชิก / เข้าสู่ระบบ / ออกจากระบบ** พัฒนาด้วย Go (Gin Framework) ใช้ JWT ในการยืนยันตัวตน และ Redis สำหรับเก็บ token ที่ใช้ session control

---

## 📦 เทคโนโลยีที่ใช้

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- [JWT (JSON Web Token)](https://jwt.io/)
- [Docker](https://www.docker.com/) (สำหรับ Redis)
- `.env` สำหรับจัดการ secret key

---

## 📁 โครงสร้างโปรเจกต์

login-backend/
├── handler/
│ ├── auth.go
│ └── handler.go
├── main.go
├── go.mod / go.sum
├── .env
└── README.md

yaml
คัดลอก
แก้ไข

---

## ⚙️ ติดตั้งและรันโปรเจกต์

### 1. Clone โปรเจกต์

```bash
git clone https://github.com/your-username/login-backend.git
cd login-backend
2. สร้าง .env
env
คัดลอก
แก้ไข
JWT_SECRET=your-super-secret-key
🔐 ค่านี้จะถูกใช้ในการ sign/verify JWT

3. ติดตั้ง dependencies
bash
คัดลอก
แก้ไข
go mod tidy
4. รัน Redis ด้วย Docker (ถ้ายังไม่มี Redis)
bash
คัดลอก
แก้ไข
docker run --name my-redis -p 6379:6379 -d redis
5. รันแอป
bash
คัดลอก
แก้ไข
go run main.go
✅ Server จะรันที่ http://localhost:8080

🛠️ API Endpoints
Method	Endpoint	Description	Auth Required
POST	/signup	สมัครสมาชิกใหม่	❌
POST	/login	เข้าสู่ระบบ	❌
POST	/api/logout	ออกจากระบบ	✅ ต้องใช้ JWT

🔑 JWT ต้องใส่ใน Authorization header แบบ:

makefile
คัดลอก
แก้ไข
Authorization: Bearer <your_token_here>
📌 ตัวอย่าง JSON Request
สมัครสมาชิก
json
คัดลอก
แก้ไข
POST /signup
{
  "username": "admin",
  "password": "123456"
}
เข้าสู่ระบบ
json
คัดลอก
แก้ไข
POST /login
{
  "username": "admin",
  "password": "123456"
}
ออกจากระบบ (ต้องแนบ JWT)
http
คัดลอก
แก้ไข
POST /api/logout
Authorization: Bearer <your_token>
🚀 ฟีเจอร์เพิ่มเติม (แนะนำในอนาคต)
ล็อคอินหลาย session

กำหนดเวลาหมดอายุ JWT

เข้ารหัส password ด้วย bcrypt

Swagger UI สำหรับ API docs

👨‍💻 ผู้พัฒนา
ชื่อผู้พัฒนา: Tuang

มหาวิทยาลัย/องค์กร: [Ubon Ratchathani University]

📝 License
Distributed under the MIT License. See LICENSE for more information.

yaml
คัดลอก
แก้ไข

---

หากคุณต้องการให้ใส่ลิงก์ GitHub จริง, ลิงก์ Swagger หรืออยากเปลี่ยนชื่อ repo — บอกได้เลยครับ ผมช่วยป
