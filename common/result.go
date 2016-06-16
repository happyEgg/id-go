package common

import (
	"github.com/astaxie/beego"
)

type JsonResult struct {
	Code   int32  `json:"code"`
	Uid    int64  `json:"uid,omitempty"`
	ErrMsg string `json:"err_msg,omitempty"`
}

func Result(c beego.Controller, err error, uid int64) {
	result := new(JsonResult)

	if err != nil {
		errResult := &JsonResult{1000, 0, "系统出错"}
		if beego.BConfig.RunMode == "dev" {
			errResult.ErrMsg = err.Error()
		}

		c.Data["json"] = errResult
		c.ServeJSON()
		return
	}

	result.Uid = uid

	c.Data["json"] = result
	c.ServeJSON()
	return
}
