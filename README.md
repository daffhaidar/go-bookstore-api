# 📚 Bookstore API

A simple REST API for bookstore management using Go. This project is created as an example of REST API implementation with Go that covers various important concepts in backend development.

## 🌟 Features

- ✨ CRUD operations for books
- 🚀 Concurrent search functionality
- 📝 Middleware for logging
- 🔄 JSON parsing and error handling
- 🔒 Mutex for thread safety
- 📖 Swagger UI documentation
- 🧪 Postman Collection for testing
- 📚 Book categorization and difficulty level

## 📋 Book Categories
- 🔧 Backend Development
- 🎨 Frontend Development
- 🛠️ DevOps
- 💾 Database
- 🧮 Algorithms & Data Structures

## 📊 Difficulty Level (Effort)
- 🟢 Beginner
- 🟡 Intermediate
- 🔴 Advanced

## 🛣️ Endpoints

- `GET /books` - Get all books
- `GET /books/{id}` - Get books by ID
- `POST /books` - Add new books
- `PUT /books/{id}` - Update books
- `DELETE /books/{id}` - Delete books
- `GET /books/search?title=:title` - Search for books by title (using concurrent processing)

## 📖 API Documentation

Swagger UI is available at: http://localhost:8080/swagger/

## 🚀 How to Run

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

4. Run the server:
   ```bash
   go run main.go
   ```
## 🧪 Testing with Postman

1. Import the `postman/Bookstore.postman_collection.json` file into Postman
2. All endpoints are configured and ready to use
3. Collection includes all CRUD and search operations

## 📚 Sample Book

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
"effort": "Beginners",
    "is_available": true
}
```
## 💡 Demonstrated Concepts

1. **REST API**
- Standard HTTP methods implementation
- JSON response handling
- URL parameter parsing

2. **Concurrency**
- Goroutines for concurrent search
- Channels for communication between goroutines
- WaitGroup for synchronization
- Mutex for thread safety

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
- Custom types for Category and Effort
- Constants for valid values
- Structured data modeling

## 🛠️ Tech Stack

- Go 1.21
- Gorilla Mux (Router)
- Swagger (API Documentation)
- Postman (API Testing)

## 📝 License

MIT License - feel free to use this project for learning purposes!
