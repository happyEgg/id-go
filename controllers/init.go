package controllers

import (
	"log"
	"os"

	"id-go/daos"
	"id-go/models"
)

var tempNum int64
var MAXSEQ int64
var orderTempNum int64
var orderMaxSeq int64

func init() {
	InitUidMaxSeq()
	InitOrderMaxSeq()
}

func InitUidMaxSeq() {
	maxSeq, err := daos.ReadUidMaxSeq()
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	MAXSEQ = maxSeq.MaxSeq + models.Step
	tempNum = maxSeq.MaxSeq
}

func InitOrderMaxSeq() {
	orderSeq, err := daos.ReadOrderMaxSeq()
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	orderMaxSeq = orderSeq.MaxSeq + models.Step
	orderTempNum = orderSeq.MaxSeq
}
