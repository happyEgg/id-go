package daos

import (
	"errors"

	"github.com/astaxie/beego/orm"

	"id-go/models"
)

func ReadMaxSeq(name string) (*models.MaxSeq, error) {
	var maxSeq = new(models.MaxSeq)

	o := models.BeegoOrm

	err := o.QueryTable("max_seq").Filter("name", name).One(maxSeq)
	if err == orm.ErrNoRows {
		err = errors.New("没有此条数据，请确定数据库max_seq表中有name字段为" + name + "的数据")
		return nil, err
	}

	_, err = o.QueryTable("max_seq").Filter("name", name).Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		return nil, err
	}

	return maxSeq, nil
}

func UpdateMaxSeq(name string) {
	o := models.BeegoOrm

	_, err := o.QueryTable("max_seq").Filter("name", name).Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		WriteLog(err.Error())
		return
	}
	return
}
