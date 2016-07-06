package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var pointMutex sync.Mutex

type PointController struct {
	beego.Controller
}

func (c *UidController) GetPointId() {
	pointMutex.Lock()
	defer pointMutex.Unlock()

	pointTempNum++

	if pointTempNum >= pointMaxSeq {
		go daos.UpdateMaxSeq("point")

		pointMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, pointTempNum)
}
