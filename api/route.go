package api

import "github.com/gin-gonic/gin"

// create a struct

type Config struct {
	Router *gin.Engine
}

// make reciever func of type struct
func (app *Config) Routes() {
	app.Router.POST("/otp", app.sendOTP())
	app.Router.POST("/verifyOtp", app.verifyOTP())
}
