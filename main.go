package main

import (
	"xitulu/model"
	"xitulu/models"
	"xitulu/router"
)

func main() {
	model.SetupDb()
	models.Setup()
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
