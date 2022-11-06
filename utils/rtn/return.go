package rtn

import "github.com/gin-gonic/gin"

// 通用数据返回
func ReturnJson(status int, data interface{}, msg interface{}) (gh gin.H) {
	gh = gin.H{
		"status":  status,
		"data":    data,
		"message": msg,
	}
	return gh
}
