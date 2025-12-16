package database

import (
	"context"
	"goApp/models/db"
	"time"
)

func CreateEmployee(emp *db.Employee) (*db.Employee, error) {

	// creating the database query
	query := `
		INSERT INTO employees (name, email, tel, age, dept, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	// uisng the time to get current time correctly
	emp.CreatedAt = time.Now()

	// querying the actual commands to the database to create an employee in it
	err := DB.QueryRow(context.Background(), query,
		emp.Name,
		emp.Email,
		emp.Tel,
		emp.Age,
		emp.Dept,
	).Scan(&emp.Id)

	// handling the err
	if err != nil {
		return nil, err
	}

	return emp, nil

}


func GetEmployeesById(id int)(*db.Employee,error){

	// setting up a data holder
	var emp db.Employee

	// setting up the query 
	query:=`SELECT name, email, tel, age, dept, created_at
			FROM Employees 
			WHERE id=$1`

	// quariying to the database
	err:=DB.QueryRow(context.Background(),query,id).
	Scan(&emp.Name,&emp.Email,&emp.Tel,&emp.Age,&emp.Dept,&emp.CreatedAt)

	// err handling
	if err!=nil{
		return nil,err
	}

	return &emp,nil


}
