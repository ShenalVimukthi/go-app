package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"goApp/services"
	"errors"
)


func GetEmployee(c *gin.Context){

	// getting the id parameter from the HTTP request
	idParam:=c.Param("id")

	// converting string to int using strconv
	convId,err:=strconv.Atoi(idParam)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":" Invalid id"})
		return
	}

	//get the user by calling the services
	emp,err:=services.GetEmployee(convId)
	if err!=nil{
		// employ not found error
		if errors.Is(err,services.ErrUserNotFound){
			c.JSON(http.StatusNotFound,err.Error())
			return
		}
		// generalized error
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Internal Server Error"})
		return
	}

	// if no errors the employee is send back as a JSON
	c.JSON(http.StatusOK,emp)


}