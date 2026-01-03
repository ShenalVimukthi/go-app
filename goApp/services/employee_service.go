package services

import (
	"database/sql"
	"goApp/database"
	"goApp/mappers"
	"goApp/models/request"
	"goApp/models/response"

	"github.com/jackc/pgconn"

	// "goApp/models/db"
	"errors"
	"fmt"
)

// create custom errors
var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserExists            = errors.New("user already exists")
	ErrInvalidEmployee       = errors.New("invalid employee data")
	ErrEmployeeAlreadyExists = errors.New("employee already exists")
	ErrInvalidReference      = errors.New("invalid related data")
	ErrMissingRequiredField  = errors.New("missing required field")
	ErrCreateEmployeeFailed  = errors.New("create employee failed")
)

// mock data

// functions to do CRUD operations on employee data
func GetEmployee(id int) (*response.EmployeeResponse, error) {

	// sending employee data inside the database querying section
	ReceivedEmp, err := database.GetEmployeesById(id)

	// handling errors
	if err != nil {
		return nil, err
	}

	// emp nil error handling
	if ReceivedEmp == nil {
		return nil, fmt.Errorf("GetEmployeesById failed for id %d: %w", id, ErrUserNotFound)
	}

	// feeding the data into response model using mapper
	emp := mappers.ToEmployeeResponse(ReceivedEmp)

	return emp, nil

}

func CreateEmployee(employee *request.CreateEmployee) (*response.EmployeeResponse, error) {

	// converting data to database model
	dbFormat := mappers.ToDataBaseModel(employee)

	// sending the employee data to table
	receivedEmp, err := database.CreateEmployee(dbFormat)

	// error handling
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505": // unique_violation
				return nil, ErrEmployeeAlreadyExists

			case "23503": // foreign_key_violation
				return nil, ErrInvalidReference

			case "23502": // not_null_violation
				return nil, ErrMissingRequiredField
			}
		}

		// fallback error
		return nil, fmt.Errorf("CreateEmployee failed: %w", ErrCreateEmployeeFailed)
	}

	// reconverting the db model to response model
	emp := mappers.ToEmployeeResponse(receivedEmp)

	return emp, nil

}

func UpdateEmployee(id int, employee *request.UpdateEmployee) (*response.EmployeeResponse, error) {

	// converting data to database model
	dbFormat := mappers.ToDataBaseModelForUpdate(employee)

	// sending the employee data to table
	receivedEmp, err := database.UpdateEmployee(id, dbFormat)

	// error handling
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505": // unique_violation
				return nil, ErrEmployeeAlreadyExists

			case "23503": // foreign_key_violation
				return nil, ErrInvalidReference

			case "23502": // not_null_violation
				return nil, ErrMissingRequiredField
			}
		}

		// fallback error
		return nil, fmt.Errorf("UpdateEmployee failed: %w", ErrCreateEmployeeFailed)
	}

	// reconverting the db model to response model
	emp := mappers.ToEmployeeResponse(receivedEmp)

	return emp, nil

}

func DeleteEmployee(id int) (*response.EmployeeResponse, error) {

	// sending employee data inside the database querying section
	ReceivedEmp, err := database.DeleteEmployee(id)

	// handling errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("DeleteEmployee failed for id %d: %w", id, ErrUserNotFound)
		}
		return nil, err
	}

	// emp nil error handling
	if ReceivedEmp == nil {
		return nil, fmt.Errorf("DeleteEmployee failed for id %d: %w", id, ErrUserNotFound)
	}

	// feeding the data into response model using mapper
	emp := mappers.ToEmployeeResponse(ReceivedEmp)

	return emp, nil

}
