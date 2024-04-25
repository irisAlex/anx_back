package response

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func DownFile(c *gin.Context, f *excelize.File) {
	filename := url.QueryEscape("A3.xlsx")

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", filename))
	c.Writer.Header().Set("Content-Type", "application/vnd.ms-excel")
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	c.File("/uploads/file/" + filename)
	// c.Writer.WriteHeader(http.StatusOK)
	// c.Header("Content-Disposition", "attachment; filename=hello.txt")
	// c.Header("Content-Type", "application/text/plain")
	// c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	// c.Writer.Write([]byte(content))
	// c.Header("Content-Type", "application/octet-stream")
	// c.Header("Content-Disposition", "attachment; filename=A3.xlsx")
	// c.Header("Content-Transfer-Encoding", "binary")
	// c.File("A3报告.xlsx")
	//Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}
