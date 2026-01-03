package request

type CreateEmployee struct{

	Name string `json:"name"  binding:"required"`
	Email string`json:"email" binding:"required,email"`
	Tel string  `json:"tel"   binding:"required"`
	Age int 	`json:"age"   binding:"required,min=18"`
	Dept string `json:"dept"  binding:"required"`

}

type UpdateEmployee struct{

	Name string `json:"name"  binding:"required"`
	Email string`json:"email" binding:"required,email"`
	Tel string  `json:"tel"   binding:"required"`
	Age int 	`json:"age"   binding:"required,min=18"`
	Dept string `json:"dept"  binding:"required"`

}