package helper

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExportCsv(t *testing.T) {
	fileName := "foo.csv"
	go func() {
		r := gin.Default()
		r.GET("/exportCsv", func(c *gin.Context) {
			ExportCsv(c.Writer, [][]string{
				[]string{"1", "2", "3"},
				[]string{"4", "5", "6"},
			}, fileName)
		})
		r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	}()
	time.Sleep(500 * time.Millisecond)
	r, err := http.Get("http://localhost:8080/exportCsv")
	assert.Nil(t, err)
	assert.Equal(t, "text/csv; charset=utf-8", r.Header.Get("Content-Type"))
	assert.Equal(t, fmt.Sprintf("attachment; filename=%s", fileName), r.Header.Get("Content-Disposition"))

}
