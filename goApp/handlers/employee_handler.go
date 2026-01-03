package handlers

import (
	"errors"
	"goApp/models/request"
	"goApp/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {

	// getting the id parameter from the HTTP request
	idParam := c.Param("id")

	// converting string to int using strconv
	convId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " Invalid id"})
		return
	}

	//get the user by calling the services
	emp, err := services.GetEmployee(convId)
	if err != nil {
		// employ not found error
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		// generalized error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// if no errors the employee is send back as a JSON
	c.JSON(http.StatusOK, emp)

}


func CreateEmployee(c *gin.Context){

	// creating the structure to store the bind data
	var emp request.CreateEmployee

	//binding the data that get from user side
	if err:=c.ShouldBindJSON(&emp); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	// sending the actual data to store in db
	savedEmp,err:=services.CreateEmployee(&emp)

		// handling generic and non generic which means datbase errors in it
		if err != nil {

		switch err {

		case services.ErrEmployeeAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})

		case services.ErrMissingRequiredField:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case services.ErrInvalidReference:
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error occurred",
			})
		}
		return
	}


	//if there are no errors returning the data updated with id
	c.JSON(http.StatusCreated,savedEmp)
	

}

func UpdateEmployee(c *gin.Context){

	// getting the id parameter from the HTTP request
	idParam := c.Param("id")

	// converting string to int using strconv
	convId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	// creating the structure to store the bind data
	var emp request.UpdateEmployee

	//binding the data that get from user side
	if err:=c.ShouldBindJSON(&emp); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	// sending the actual data to update in db
	updatedEmp,err:=services.UpdateEmployee(convId,&emp)

		// handling generic and non generic which means datbase errors in it
		if err != nil {

		switch err {

		case services.ErrEmployeeAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})

		case services.ErrMissingRequiredField:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case services.ErrInvalidReference:
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error occurred",
			})
		}
		return
	}


	//if there are no errors returning the updated data
	c.JSON(http.StatusOK,updatedEmp)
	

}

func DeleteEmployee(c *gin.Context){

	// getting the id parameter from the HTTP request
	idParam := c.Param("id")

	// converting string to int using strconv
	convId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	//delete the user by calling the services
	emp, err := services.DeleteEmployee(convId)
	if err != nil {
		// employ not found error
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		// generalized error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// if no errors the employee is send back as a JSON
	c.JSON(http.StatusOK, emp)

}
