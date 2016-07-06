package controllers

import (
	"log"
	"os"

	"id-go/daos"
	"id-go/models"
)

var uidTempNum, uidMaxSeq int64
var orderTempNum, orderMaxSeq int64
var albumMaxSeq, albumTempNum int64
var blockMaxSeq, blockTempNum int64
var likeMaxSeq, likeTempNum int64
var visitorMaxSeq, visitorTempNum int64
var userGiftMaxSeq, userGiftTempNum int64
var sendGiftMaxSeq, sendGiftTempNum int64
var accountOptMaxSeq, accountOptTempNum int64
var pointMaxSeq, pointTempNum int64

func init() {
	InitUidMaxSeq()
	InitOrderMaxSeq()
	InitAlbumMaxSeq()
	InitBlockMaxSeq()
	InitLikeMaxSeq()
	InitVisitorMaxSeq()
	InitUserGiftMaxSeq()
	InitSendGiftMaxSeq()
	InitAccountOptMaxSeq()
	InitPointMaxSeq()
}

func InitUidMaxSeq() {
	maxSeq, err := daos.ReadMaxSeq("user")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	uidMaxSeq = maxSeq.MaxSeq + models.Step
	uidTempNum = maxSeq.MaxSeq
}

func InitOrderMaxSeq() {
	orderSeq, err := daos.ReadMaxSeq("order")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	orderMaxSeq = orderSeq.MaxSeq + models.Step
	orderTempNum = orderSeq.MaxSeq
}

func InitAlbumMaxSeq() {
	seq, err := daos.ReadMaxSeq("album")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	albumMaxSeq = seq.MaxSeq + models.Step
	albumTempNum = seq.MaxSeq
}

func InitBlockMaxSeq() {
	seq, err := daos.ReadMaxSeq("block")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	blockMaxSeq = seq.MaxSeq + models.Step
	blockTempNum = seq.MaxSeq
}

func InitLikeMaxSeq() {
	seq, err := daos.ReadMaxSeq("like")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	likeMaxSeq = seq.MaxSeq + models.Step
	likeTempNum = seq.MaxSeq
}

func InitVisitorMaxSeq() {
	seq, err := daos.ReadMaxSeq("visitor")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	visitorMaxSeq = seq.MaxSeq + models.Step
	visitorTempNum = seq.MaxSeq
}

func InitUserGiftMaxSeq() {
	seq, err := daos.ReadMaxSeq("user_gift")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	userGiftMaxSeq = seq.MaxSeq + models.Step
	userGiftTempNum = seq.MaxSeq
}

func InitSendGiftMaxSeq() {
	seq, err := daos.ReadMaxSeq("send_gift")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	sendGiftMaxSeq = seq.MaxSeq + models.Step
	sendGiftTempNum = seq.MaxSeq
}

func InitAccountOptMaxSeq() {
	seq, err := daos.ReadMaxSeq("account_opt")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	accountOptMaxSeq = seq.MaxSeq + models.Step
	accountOptTempNum = seq.MaxSeq
}

func InitPointMaxSeq() {
	seq, err := daos.ReadMaxSeq("point")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	pointMaxSeq = seq.MaxSeq + models.Step
	pointTempNum = seq.MaxSeq
}
