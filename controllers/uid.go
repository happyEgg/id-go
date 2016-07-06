package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var uidMutex sync.Mutex

type UidController struct {
	beego.Controller
}

func (c *UidController) GetUid() {
	uidMutex.Lock()
	defer uidMutex.Unlock()

	uidTempNum++

	if uidTempNum >= uidMaxSeq {
		go daos.UpdateMaxSeq("user")

		uidMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, uidTempNum)
}
