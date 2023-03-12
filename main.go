package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thelearningbard/basic-server/controllers"
)

func main() {
	router := gin.Default() // Address: localhost:8080

	router.POST("/person", controllers.CreatePerson)
	router.GET("/person/:id", controllers.GetPerson)
	router.DELETE("/person/:id", controllers.DeletePerson)
	router.PUT("/person/:id", controllers.UpdatePerson)

	router.Run()
}
