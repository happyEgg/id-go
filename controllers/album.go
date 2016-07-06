package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var albumMutex sync.Mutex

type UidController struct {
	beego.Controller
}

func (c *UidController) GetUid() {
	albumMutex.Lock()
	defer albumMutex.Unlock()

	albumTempNum++

	if albumTempNum >= albumMaxSeq {
		go daos.UpdateOrderMaxSeq()

		albumMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, albumTempNum)
}
