package models

type MaxSeq struct {
	Id     int32
	Name   string `orm:"unique,size(50)"`
	MaxSeq int64
}
