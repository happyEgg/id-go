package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var accountOptMutex sync.Mutex

type AccountOptController struct {
	beego.Controller
}

func (c *UidController) GetAccountOptId() {
	accountOptMutex.Lock()
	defer accountOptMutex.Unlock()

	accountOptTempNum++

	if accountOptTempNum >= accountOptMaxSeq {
		go daos.UpdateMaxSeq("account_opt")

		accountOptMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, accountOptTempNum)
}
