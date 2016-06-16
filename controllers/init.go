package controllers

import (
	"log"
	"os"

	"id-go/daos"
	"id-go/models"
)

var tempNum int64
var MAXSEQ int64

func init() {
	InitMaxSeq()
}

func InitMaxSeq() {
	maxSeq, err := daos.ReadMaxSeq()
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}

	MAXSEQ = maxSeq.MaxSeq + models.Step
	tempNum = maxSeq.MaxSeq
}
