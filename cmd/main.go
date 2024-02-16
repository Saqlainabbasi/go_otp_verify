package main

import (
	"github.com/Saqlainabbasi/go_otp_verify/api"
	"github.com/gin-gonic/gin"
)

func main() {
	//create a router with gin
	router := gin.Default()
	//initialize the routes config
	app := api.Config{Router: router}
	// Use the router
	app.Routes()
	//start the server on port 8080
	router.Run(":8080")
}
