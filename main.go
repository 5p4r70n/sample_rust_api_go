package main

import (
	"github.com/gin-gonic/gin"
	"restapi/controllers"
)

func main() {

	router := gin.Default()
	
	router.GET("/person/:person_id/info", controllers.GetPersonById)
	router.POST("/person/create", controllers.InsertData)

	router.Run("0.0.0.0:8080")

}