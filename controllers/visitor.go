package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var visitorMutex sync.Mutex

type VisitorController struct {
	beego.Controller
}

func (c *VisitorController) GetVisitorId() {
	visitorMutex.Lock()
	defer visitorMutex.Unlock()

	visitorTempNum++

	if visitorTempNum >= visitorMaxSeq {
		go daos.UpdateMaxSeq("visitor")

		visitorMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, visitorTempNum)
}
