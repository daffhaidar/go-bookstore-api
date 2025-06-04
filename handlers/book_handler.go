package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"bookstore/models"

	"github.com/gorilla/mux"
)

var (
	books = make(map[string]models.Book)
	mutex = &sync.RWMutex{}
)

func init() {
	// Initialize with sample books
	sampleBooks := []models.Book{
		{
			ID:          "6",
			Title:       "Go Programming: Zero to Hero",
			Author:      "Mat Ryer",
			Price:       35.99,
			Category:    models.BackEnd,
			Effort:      models.Beginner,
			IsAvailable: true,
		},
		{
			ID:          "7",
			Title:       "Vue.js 3 By Example",
			Author:      "John Au-Yeung",
			Price:       42.99,
			Category:    models.FrontEnd,
			Effort:      models.Beginner,
			IsAvailable: true,
		},
		{
			ID:          "8",
			Title:       "MongoDB: The Definitive Guide",
			Author:      "Shannon Bradshaw",
			Price:       49.99,
			Category:    models.Database,
			Effort:      models.Intermediate,
			IsAvailable: true,
		},
		{
			ID:          "9",
			Title:       "Kubernetes in Action",
			Author:      "Marko Luksa",
			Price:       55.99,
			Category:    models.DevOps,
			Effort:      models.Intermediate,
			IsAvailable: true,
		},
	}

	for _, book := range sampleBooks {
		books[book.ID] = book
	}
}

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	booksList := make([]models.Book, 0, len(books))
	for _, book := range books {
		booksList = append(booksList, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksList)
}

// GetBook godoc
// @Summary Get a book by ID
// @Description Get details of a specific book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	mutex.RLock()
	book, found := books[params["id"]]
	mutex.RUnlock()

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	json.NewEncoder(w).Encode(book)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the store
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book object"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	mutex.Lock()
	books[book.ID] = book
	mutex.Unlock()

	json.NewEncoder(w).Encode(book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update an existing book's details
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Book object"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, found := books[params["id"]]; !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	book.ID = params["id"]
	books[params["id"]] = book
	json.NewEncoder(w).Encode(book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book from the store
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	mutex.Lock()
	defer mutex.Unlock()

	if _, found := books[params["id"]]; !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	delete(books, params["id"])
	w.WriteHeader(http.StatusNoContent)
}

// SearchBooks godoc
// @Summary Search books by title
// @Description Search books using concurrent processing
// @Tags books
// @Accept json
// @Produce json
// @Param title query string true "Book title to search"
// @Success 200 {array} models.Book
// @Router /books/search [get]
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("title")

	mutex.RLock()
	booksCopy := make([]models.Book, 0, len(books))
	for _, book := range books {
		booksCopy = append(booksCopy, book)
	}
	mutex.RUnlock()

	resultChan := make(chan models.Book)
	var wg sync.WaitGroup

	for _, book := range booksCopy {
		wg.Add(1)
		go func(b models.Book) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			if b.Title == query {
				resultChan <- b
			}
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := []models.Book{}
	for book := range resultChan {
		results = append(results, book)
	}

	json.NewEncoder(w).Encode(results)
} 