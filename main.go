package main

import (
	"log"
	"xitulu/models"
	"xitulu/router"
)

func main() {
	models.Setup()
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		log.Println("err", err)
	}
}
