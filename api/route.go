package api

import "github.com/gin-gonic/gin"

// create a struct
// it has a field
type Config struct {
	Router *gin.Engine
}

// make reciever func
func (app *Config) Routes() {
	app.Router.POST("/otp")
	app.Router.POST("/verifyOtp")
}
