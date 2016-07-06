package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var likeMutex sync.Mutex

type LikeController struct {
	beego.Controller
}

func (c *UidController) GetLikeId() {
	likeMutex.Lock()
	defer likeMutex.Unlock()

	likeTempNum++

	if likeTempNum >= likeMaxSeq {
		go daos.UpdateMaxSeq("like")

		likeMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, likeTempNum)
}
