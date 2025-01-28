package models

import (
	"time"
)

// create post model 
type Post struct {
	ID			string		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	Date		time.Time	`json:"date"`
	Category	*Category	`json:"category"`
}

