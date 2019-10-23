package util

import (
	"VueElementAdminGoBackend/pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, j interface{}) int {
	err := c.ShouldBind(j)

	if err != nil {
		return e.WarningInvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(j)

	if err != nil {
		fmt.Println(j, err)
		return e.Error
	}

	if !check {
		MarkErrors(valid.Errors)
		return e.WarningInvalidParams
	}

	return e.Success
}
