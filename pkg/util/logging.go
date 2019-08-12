package util

import (
	"github.com/astaxie/beego/validation"
	"log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Println(err.Key, err.Message)  // 打印日志
	}

	return
}