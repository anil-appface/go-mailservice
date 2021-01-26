package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/anil-appface/go-mailservice/resthandlers"
)

func main() {

	port := ":8000"

	//set gin as release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		emailHandler := resthandlers.NewMailHandler()
		v1.POST("sendEmail", emailHandler.SendEmail)
	}

	// listen and serve on 8000
	router.Run(port)
}
