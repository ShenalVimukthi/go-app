package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/sendEmp",func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"employees sending",
		})
	})
	r.Run(":8500")
}