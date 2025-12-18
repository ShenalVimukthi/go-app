package handlers

import (
	"errors"
	"goApp/services"
	"net/http"
	"strconv"
	"goApp/models/request"
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
	if err!=nil{
		if errors.Is(err,services.ErrCreateEmployeeFailed){
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error occured"})
		return
	}

	//if there are no errors returning the data updated with id
	c.JSON(http.StatusCreated,savedEmp)

}
