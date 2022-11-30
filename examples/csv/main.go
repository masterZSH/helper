package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/helper"
)

func main() {
	r := gin.Default()
	r.GET("/exportCsv", func(c *gin.Context) {
		helper.ExportCsv(c.Writer, [][]string{
			[]string{"1", "2", "3"},
			[]string{"4", "5", "6"},
		}, "1.csv")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	// http://locahost:8080/exportCsv
}
