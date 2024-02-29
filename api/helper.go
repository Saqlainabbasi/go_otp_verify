package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// new intance of validator.....
var validate = validator.New()

// reciever func of type Config
// takes gin context and data as parameter, return error
func (app *Config) validateBody(ctx *gin.Context, data any) error {
	//validate the req body
	if err := ctx.BindJSON(&data); err != nil {
		return err
	}
	if err := validate.Struct(&data); err != nil {
		return err
	}
	return nil
}

// reciever func to WriteJson
func (app *Config) writeJSON(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, jsonResponse{Status: status, Message: "Success", Data: data})
}

// reciever func
// takes three parameter, the status parameter considered as default param
// as go does not support the default parameter we can achive this with using the empty struct
func (app *Config) errorJSON(ctx *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	ctx.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}
