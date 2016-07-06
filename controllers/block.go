package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var blockMutex sync.Mutex

type BlockController struct {
	beego.Controller
}

func (c *UidController) GetBlockId() {
	blockMutex.Lock()
	defer blockMutex.Unlock()

	blockTempNum++

	if blockTempNum >= blockMaxSeq {
		go daos.UpdateMaxSeq("block")

		blockMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, blockTempNum)
}
