package util

import (
	"VueElementAdminGoBackend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"ok"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    code,
		Message: e.GetMsg(code),
		Data:    e.GetData(code, data),
	})
	return
}
