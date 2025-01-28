package models

// create category model 
type Category struct {
	ID		string	`json:"id"`
	Post	*Post	`json:"post"`
}