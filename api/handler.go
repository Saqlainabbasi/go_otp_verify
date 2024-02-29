package api

import (
	"context"
	"time"

	"github.com/Saqlainabbasi/go_otp_verify/data"
	"github.com/gin-gonic/gin"
)

// timeout for req variable
var appTimeout = time.Second * 10

// func to send otp
func (app *Config) sendOTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancle := context.WithTimeout(context.Background(), appTimeout) //create the context
		var payload data.OTPData
		defer cancle()

		app.validateBody(ctx, &payload)

		//validate the data
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		//call the service func
		_, err := app.TwilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}
		//send resp back
		app.writeJSON(ctx, 200, "OTP sent successfully")
	}
}

// func to verify OTP
func (app *Config) verifyOTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//create new context
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		//create payload for verify req
		var payload data.VerifyData
		//cancel the context
		defer cancel()
		app.validateBody(ctx, &payload)

		//make the data for the req
		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		//call the service
		err := app.TwilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		//check err
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}
		//send resp
		app.writeJSON(ctx, 200, "Valid Otp")
	}
}
