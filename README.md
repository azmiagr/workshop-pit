# Workshop PIT - Book Management API

Aplikasi REST API sederhana untuk manajemen buku menggunakan Go, Gin Framework, dan GORM dengan database MariaDB/MySQL.

## 📋 Daftar Isi

- [Fitur](#fitur)
- [Teknologi](#teknologi)
- [Struktur Folder](#struktur-folder)
- [Cara Menjalankan](#cara-menjalankan)
- [API Endpoints](#api-endpoints)
- [Contoh Request](#contoh-request)

## ✨ Fitur

- CRUD (Create, Read, Update, Delete) untuk manajemen buku
- Clean Architecture dengan pemisahan layer (Repository, Service, Handler)
- Auto migration database menggunakan GORM
- Response API yang konsisten
- Validasi input menggunakan Gin binding

## 🛠 Teknologi

- **Go** 1.24.1
- **Gin** - HTTP Web Framework
- **GORM** - ORM untuk Go
- **MariaDB/MySQL** - Database
- **godotenv** - Environment variable management
- **UUID** - Unique identifier generation

## 📁 Struktur Folder
```
workshop-pit/
├── cmd/
│ └── app/
│ └── main.go # Entry point aplikasi
├── entity/ # Database entities (GORM models)
│ ├── book.go
│ ├── loan.go
│ └── user.go
├── internal/
│ ├── handler/
│ │ └── rest/ # REST API handlers (Controller layer)
│ │ ├── rest.go # Router setup dan server initialization
│ │ └── book.go # Book endpoints handlers
│ ├── repository/ # Repository layer (Data access)
│ │ ├── repository.go # Repository aggregator
│ │ └── book.go # Book repository implementation
│ └── service/ # Service layer (Business logic)
│ ├── service.go # Service aggregator
│ └── book.go # Book service implementation
├── model/ # Request/Response DTOs
│ └── book.go
├── pkg/
│ ├── config/ # Configuration management
│ │ ├── config.go # Environment loader
│ │ └── database.go # Database connection string builder
│ ├── database/
│ │ └── mariadb/ # Database connection & migration
│ │ ├── mariadb.go
│ │ └── migrate.go
│ └── response/ # Standard API response format
│ └── response.go
├── .env # Environment variables (jangan commit!)
├── .env.example # Template environment variables
├── go.mod # Go module dependencies
└── README.md
```

### Penjelasan Layer

#### 1. **Repository Layer** (`internal/repository/`)
- Bertanggung jawab untuk akses data ke database
- Berinteraksi langsung dengan GORM
- Melakukan operasi CRUD dasar
- Interface-based untuk memudahkan testing dan maintainability

#### 2. **Service Layer** (`internal/service/`)
- Berisi business logic aplikasi
- Melakukan validasi dan transformasi data
- Mengatur transaction database
- Memanggil repository untuk operasi data
- Mengkonversi antara Entity dan Model (DTO)

#### 3. **Handler/REST Layer** (`internal/handler/rest/`)
- Menerima HTTP request
- Melakukan validasi input dari request
- Memanggil service untuk memproses business logic
- Mengembalikan HTTP response dengan format yang konsisten

## 🚀 Cara Menjalankan

### Prerequisites

- Go 1.24.1 atau lebih tinggi
- MariaDB atau MySQL
- Git

### Langkah-langkah

1. **Clone repository**
git clone <repository-url>
```
cd workshop-pit
```

2. **Install dependencies**
```
go mod download
```

3. **Setup environment variables**   
Copy file `.env.example` menjadi `.env`:
```
cp .env.example .env
```
Kemudian edit file `.env` sesuai dengan konfigurasi database Anda:

```    
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=workshop_pit

ADDRESS=localhost
PORT=8080
```
 4. **Buat database**   
Buat database baru di MariaDB/MySQL:
```
CREATE DATABASE workshop_pit;
```

5. **Jalankan aplikasi**
```
go run cmd/app/main.go
```
Aplikasi akan berjalan di `http://localhost:8080` (sesuai konfigurasi PORT di .env). 
Auto migration akan dijalankan otomatis saat aplikasi start.

6. **Test API**
Gunakan Postman, cURL, atau tools lainnya untuk test API endpoints.

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1`

### Book Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/book/get-all-books` | Mendapatkan semua buku |
| GET | `/book/get-book-by-id/:id` | Mendapatkan buku berdasarkan ID |
| POST | `/book/create-book` | Membuat buku baru |
| PUT | `/book/update-book/:id` | Update buku berdasarkan ID |
| DELETE | `/book/delete-book/:id` | Hapus buku berdasarkan ID |

## 📝 Contoh Request

### 1. Create Book
**POST** `/api/v1/book/create-book`

**Request Body:**
```
{
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "publisher": "Prentice Hall",
  "year": 2008,
  "isbn": "978-0132350884",
  "stock": 10,
  "description": "A Handbook of Agile Software Craftsmanship"
}
```
**Response (201 Created):**
```
{
  "status": "success",
  "message": "Book created successfully",
  "data": {
    "book_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "publisher": "Prentice Hall",
    "year": 2008,
    "isbn": "978-0132350884",
    "stock": 10,
    "description": "A Handbook of Agile Software Craftsmanship",
    "created_at": "2025-11-08T10:30:00Z",
    "updated_at": "2025-11-08T10:30:00Z"
  }
}
```

### 2. Get All Books
**GET** `/api/v1/book/get-all-books`

**Response (200 OK):**
```
{
  "status": "success",
  "message": "Books fetched successfully",
  "data": [
    {
      "book_id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Clean Code",
      "author": "Robert C. Martin",
      "publisher": "Prentice Hall",
      "year": 2008
    }
  ]
}
```

### 3. Get Book By ID
**GET** `/api/v1/book/get-book-by-id/550e8400-e29b-41d4-a716-446655440000`

**Response (200 OK):**
```
{
  "status": "success",
  "message": "Book fetched successfully",
  "data": {
    "book_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "publisher": "Prentice Hall",
    "year": 2008,
    "isbn": "978-0132350884",
    "stock": 10,
    "description": "A Handbook of Agile Software Craftsmanship"
  }
}
```

### 4. Update Book
**PUT** `/api/v1/book/update-book/550e8400-e29b-41d4-a716-446655440000`

**Request Body:**
```
{
  "title": "Clean Code - Updated",
  "author": "Robert C. Martin",
  "publisher": "Prentice Hall",
  "year": 2008,
  "isbn": "978-0132350884",
  "stock": 15,
  "description": "A Handbook of Agile Software Craftsmanship - Updated Edition"
}
```

**Response (200 OK):**
```
{
  "status": "success",
  "message": "Book updated successfully",
  "data": {
    "book_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Clean Code - Updated",
    "author": "Robert C. Martin",
    "publisher": "Prentice Hall",
    "year": 2008,
    "isbn": "978-0132350884",
    "stock": 15,
    "description": "A Handbook of Agile Software Craftsmanship - Updated Edition",
    "created_at": "2025-11-08T10:30:00Z",
    "updated_at": "2025-11-08T11:00:00Z"
  }
}
```

### 5. Delete Book
**DELETE** `/api/v1/book/delete-book/550e8400-e29b-41d4-a716-446655440000`

**Response (200 OK):**
```
{
  "status": "success",
  "message": "Book deleted successfully",
  "data": null
}
```

### Error Response
**Response (400/500):**
```
{
  "status": "error",
  "message": "Failed to create book",
  "error": "error details here"
}
```

## 🔧 Development

### Build aplikasi
```
go build -o bin/app cmd/app/main.go
```

### Run binary
```
./bin/app
```

## 📄 License
This project is for educational purposes (Workshop PIT).

## 👥 Contributors

- Azmi Al Ghifari Rahman

---
Dibuat dengan ❤️ untuk Workshop PIT