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
		emp.CreatedAt,
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
	query:=`SELECT id, name, email, tel, age, dept, created_at
			FROM employees 
			WHERE id=$1`

	// quariying to the database
	err:=DB.QueryRow(context.Background(),query,id).
	Scan(&emp.Id,&emp.Name,&emp.Email,&emp.Tel,&emp.Age,&emp.Dept,&emp.CreatedAt)

	// err handling
	if err!=nil{
		return nil,err
	}

	return &emp,nil


}


func UpdateEmployee(id int, emp *db.Employee)(*db.Employee,error){

	// query to update employees
	query := `
		UPDATE employees
		SET name=$1, email=$2, tel=$3, age=$4, dept=$5
		WHERE id=$6
		RETURNING id, created_at`

	err:=DB.QueryRow(context.Background(),query,
		emp.Name,
		emp.Email,
		emp.Tel,
		emp.Age,
		emp.Dept,
		id,		
	).Scan(&emp.Id,&emp.CreatedAt)

	if err!=nil{
		return nil ,err
	}

	return emp ,nil

	
}


func DeleteEmployee(id int)(*db.Employee,error){

	// createing a struct of employee to return
	var emp db.Employee

	// delete query for the employee
	query:=`DELETE *
			FROM employees
			WHERE id=$1
			RETURNING id, name, email, tel, age, dept, created_at`

	err:=DB.QueryRow(context.Background(),query,id).
	Scan(&emp.Id,&emp.Name,&emp.Email,&emp.Tel,&emp.Age,&emp.Dept,&emp.CreatedAt)

	if err!=nil{
		return nil,err
	}

	return &emp,nil

}