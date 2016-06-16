package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var mutex sync.Mutex

type UidController struct {
	beego.Controller
}

func (c *UidController) GetUid() {
	mutex.Lock()
	defer mutex.Unlock()

	tempNum++

	if tempNum >= MAXSEQ {
		go daos.UpdateMaxSeq()

		MAXSEQ += models.Step
	}

	common.Result(c.Controller, nil, tempNum)
}
