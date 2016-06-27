package daos

import (
	"errors"

	"github.com/astaxie/beego/orm"

	"id-go/models"
)

func ReadOrderMaxSeq() (*models.OrderMaxseq, error) {
	var orderMax = new(models.OrderMaxseq)

	o := models.BeegoOrm

	err := o.QueryTable("order_maxseq").Filter("name", "order").One(orderMax)
	if err == orm.ErrNoRows {
		err = errors.New("没有此条数据")
		return nil, err
	}

	_, err = o.QueryTable("order_maxseq").Filter("name", "order").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		return nil, err
	}

	return orderMax, nil
}

func UpdateOrderMaxSeq() {
	o := models.BeegoOrm

	_, err := o.QueryTable("order_maxseq").Filter("name", "order").Update(orm.Params{
		"max_seq": orm.ColValue(orm.ColAdd, models.Step),
	})
	if err != nil {
		WriteLog(err.Error())
		return
	}
	return
}
