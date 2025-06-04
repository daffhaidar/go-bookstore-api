# 📚 Bookstore API

Sebuah REST API sederhana untuk manajemen toko buku menggunakan Go. Project ini dibuat sebagai contoh implementasi REST API dengan Go yang mencakup berbagai konsep penting dalam pengembangan backend.

## 🌟 Fitur

- ✨ CRUD operations untuk buku
- 🚀 Concurrent search functionality
- 📝 Middleware untuk logging
- 🔄 JSON parsing dan error handling
- 🔒 Mutex untuk thread safety
- 📖 Swagger UI documentation
- 🧪 Postman Collection untuk testing
- 📚 Kategorisasi buku dan tingkat kesulitan

## 📋 Kategori Buku
- 🔧 Backend Development
- 🎨 Frontend Development
- 🛠️ DevOps
- 💾 Database
- 🧮 Algorithms & Data Structures

## 📊 Tingkat Kesulitan (Effort)
- 🟢 Beginner
- 🟡 Intermediate
- 🔴 Advanced

## 🛣️ Endpoints

- `GET /books` - Mendapatkan semua buku
- `GET /books/{id}` - Mendapatkan buku berdasarkan ID
- `POST /books` - Menambah buku baru
- `PUT /books/{id}` - Mengupdate buku
- `DELETE /books/{id}` - Menghapus buku
- `GET /books/search?title=:title` - Mencari buku berdasarkan judul (menggunakan concurrent processing)

## 📖 Dokumentasi API

Swagger UI tersedia di: http://localhost:8080/swagger/

## 🚀 Cara Menjalankan

1. Clone repository:
   ```bash
   git clone https://github.com/username/bookstore-api.git
   cd bookstore-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Generate Swagger docs:
   ```bash
   swag init
   ```

4. Jalankan server:
   ```bash
   go run main.go
   ```

## 🧪 Testing dengan Postman

1. Import file `postman/Bookstore.postman_collection.json` ke Postman
2. Semua endpoint sudah dikonfigurasi dan siap digunakan
3. Collection mencakup semua operasi CRUD dan search

## 📚 Contoh Buku

1. Backend Development:
```json
{
    "id": "6",
    "title": "Go Programming: Zero to Hero",
    "author": "Mat Ryer",
    "price": 35.99,
    "category": "Backend Development",
    "effort": "Beginner",
    "is_available": true
}
```

2. Frontend Development:
```json
{
    "id": "7",
    "title": "Vue.js 3 By Example",
    "author": "John Au-Yeung",
    "price": 42.99,
    "category": "Frontend Development",
    "effort": "Beginner",
    "is_available": true
}
```

## 💡 Konsep yang Didemonstrasikan

1. **REST API**
   - Implementasi standar HTTP methods
   - JSON response handling
   - URL parameter parsing

2. **Concurrency**
   - Goroutines untuk concurrent search
   - Channel untuk komunikasi antar goroutines
   - WaitGroup untuk sinkronisasi
   - Mutex untuk thread safety

3. **Error Handling**
   - HTTP status codes
   - JSON error responses
   - Graceful error handling

4. **Middleware**
   - Request logging
   - Execution time tracking

5. **Documentation**
   - Swagger UI integration
   - API documentation
   - Postman collection

6. **Domain Modeling**
   - Custom types untuk Category dan Effort
   - Konstanta untuk valid values
   - Structured data modeling

## 🛠️ Tech Stack

- Go 1.21
- Gorilla Mux (Router)
- Swagger (API Documentation)
- Postman (API Testing)

## 📝 License

MIT License - feel free to use this project for learning purposes! 