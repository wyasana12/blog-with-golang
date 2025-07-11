# Blog App Berbasis Rest API dan Golang

Ini merupakan proyek berbasis Resful API untuk platform blog yang dibangun dengan menggunakan bahasa pemrograman Go dan framework Echo. Proyek ini mencakup fitur-fitur seperti Autentikasi pengguna dengan JWT, Manejemen role dengan hak akses, dan operasi CRUD untuk postingan Blog.

## Fitur

## Teknologi Yang Digunakan
- Go (Golang).
- Echo v4 (Framework).
- Postgres SQL.
- GORM.
- golang-jwt/jwt.
- bcrypt.
- godotenv.
- validator/v10.
- air.

## Instalasi dan Konfigurasi
#### 1. Kloning Repositori
```bash
git clone https://github.com/wyasana12/blog-with-golang.git
cd blog-with-golang
```

#### 2. Konfigurasi ENV
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=YOUR_PASSWORD
DB_NAME=blog

JWT_SECRET=YOUR_SECRET_KEY

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_EMAIL=YOUR_EMAIL
SMTP_PASSWORD=YOUR_PASSWORD_APP
```

#### 3. Instalasi Despendensi
```bash
go mod tidy
```

#### 4. Jalankan Aplikasi
```bash
go run ./cmd/main.go
```

## Dokumentasi API
Setelah aplikasi berhasil berjalan, dokumentasi API yang interaktif dapat diakses melalui Swagger UI di alamat:
http://localhost:8080/swagger