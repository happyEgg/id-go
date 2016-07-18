package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var reportMutex sync.Mutex

type ReportController struct {
	beego.Controller
}

func (c *ReportController) GetReportId() {
	reportMutex.Lock()
	defer reportMutex.Unlock()

	reportTempNum++

	if reportTempNum >= reportMaxSeq {
		go daos.UpdateMaxSeq("report")

		reportMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, reportTempNum)
}
