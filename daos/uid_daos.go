package daos

import (
	"errors"

	"github.com/astaxie/beego/orm"

	"id-go/models"
)

func ReadUidMaxSeq() (*models.MaxSeq, error) {
	var maxSeq = new(models.MaxSeq)

	o := models.BeegoOrm

	err := o.QueryTable("max_seq").Filter("name", "user").One(maxSeq)
	if err == orm.ErrNoRows {
		err = errors.New("没有此条数据")
		return nil, err
	}

	_, err = o.QueryTable("max_seq").Filter("name", "user").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		return nil, err
	}

	return maxSeq, nil
}

func UpdateUidMaxSeq() {
	o := models.BeegoOrm

	_, err := o.QueryTable("max_seq").Filter("name", "user").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		WriteLog(err.Error())
		return
	}
	return
}
