package db

import "time"

type Employee struct{

	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string`json:"email"`
	Tel string  `json:"tel"`
	Age int 	`json:"age"`
	Dept string `json:"dept"`
	CreatedAt time.Time `json:"created_at"`
	
}
