package daos

import (
	"errors"

	"github.com/astaxie/beego/orm"

	"id-go/models"
)

func ReadMaxSeq() (*models.MaxSeqInfo, error) {
	var maxSeq = new(models.MaxSeqInfo)

	o := models.BeegoOrm

	err := o.QueryTable("max_seq_info").Filter("table_name", "user").One(maxSeq)
	if err == orm.ErrNoRows {
		err = errors.New("没有此条数据")
		return nil, err
	}

	_, err = o.QueryTable("max_seq_info").Filter("table_name", "user").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		return nil, err
	}

	return maxSeq, nil
}

func UpdateMaxSeq() {
	o := models.BeegoOrm

	_, err := o.QueryTable("max_seq_info").Filter("table_name", "user").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		WriteLog(err.Error())
		return
	}
	return
}
