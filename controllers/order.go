package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var orderMutex sync.Mutex

type OrderController struct {
	beego.Controller
}

func (c *OrderController) GetOrderId() {
	orderMutex.Lock()
	defer orderMutex.Unlock()

	orderTempNum++

	if orderTempNum >= orderMaxSeq {
		go daos.UpdateOrderMaxSeq()

		orderMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, orderTempNum)
}
