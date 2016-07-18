package controllers

import (
	"sync"

	"github.com/astaxie/beego"

	"id-go/common"
	"id-go/daos"
	"id-go/models"
)

var albumMutex sync.Mutex

type AlbumController struct {
	beego.Controller
}

func (c *AlbumController) GetAlbumId() {
	albumMutex.Lock()
	defer albumMutex.Unlock()

	albumTempNum++

	if albumTempNum >= albumMaxSeq {
		go daos.UpdateMaxSeq("album")

		albumMaxSeq += models.Step
	}

	common.Result(c.Controller, nil, albumTempNum)
}
