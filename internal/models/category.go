package models

// create category model 
type Category struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
	Post	*Post	`json:"post"`
}