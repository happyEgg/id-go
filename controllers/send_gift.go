package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var sendGiftMutex sync.Mutex

type SendGiftController struct {
	beego.Controller
}

func (c *SendGiftController) GetSendGiftId() {
	sendGiftMutex.Lock()
	defer sendGiftMutex.Unlock()

	sendGiftTempNum++

	if sendGiftTempNum >= sendGiftMaxSeq {
		go daos.UpdateMaxSeq("send_gift")

		sendGiftMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, sendGiftTempNum)
}
