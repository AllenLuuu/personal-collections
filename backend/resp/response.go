package resp

import (
	"github.com/gin-gonic/gin"
)

const CTX_RESPONSE = "resp"

type JsonResp struct {
	ErrCode int         `json:"code"`
	ErrMsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func set_resp(c *gin.Context, resp JsonResp) {
	c.Set(CTX_RESPONSE, resp)
}

func JSON(c *gin.Context, data interface{}) {
	set_resp(c, JsonResp{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    data,
	})
}

func ERR(c *gin.Context, errcode int, errmsg string, data ...interface{}) {
	var data_item any
	if len(data) != 0 {
		data_item = data[0]
	}
	set_resp(c, JsonResp{
		ErrCode: errcode,
		ErrMsg:  errmsg,
		Data:    data_item,
	})
}
