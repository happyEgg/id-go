package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var userGiftMutex sync.Mutex

type UserGiftController struct {
	beego.Controller
}

func (c *UserGiftController) GetUserGiftId() {
	userGiftMutex.Lock()
	defer userGiftMutex.Unlock()

	userGiftTempNum++

	if userGiftTempNum >= userGiftMaxSeq {
		go daos.UpdateMaxSeq("user_gift")

		userGiftMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, userGiftTempNum)
}
