package models

type Category string

const (
	BackEnd    Category = "Backend Development"
	FrontEnd   Category = "Frontend Development"
	DevOps     Category = "DevOps"
	Database   Category = "Database"
	Algorithms Category = "Algorithms & Data Structures"
)

type Effort string

const (
	Beginner     Effort = "Beginner"
	Intermediate Effort = "Intermediate"
	Advanced     Effort = "Advanced"
)

type Book struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Price       float64  `json:"price"`
	Category    Category `json:"category"`
	Effort      Effort   `json:"effort"`
	IsAvailable bool     `json:"is_available"`
} 