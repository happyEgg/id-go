package models

type MaxSeqInfo struct {
	Id        int32
	TableName string `orm:"unique,size(50)"`
	MaxSeq    int64
}
