package util

import (
	"github.com/DeluxeYang/GinProject/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code 		int         `json:"code"`
	Message  	string      `json:"message"`
	Data 		interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: 		code,
		Message:  	e.GetMsg(code),
		Data: 		e.GetData(code, data),
	})
	return
}