package controllers

import (
	"github.com/astaxie/beego"
	"ihome/models"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController)RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

/*
 * handle area request
 */
func (c *AreaController)GetArea() {
	beego.Info("------------------connect success--------------------")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	// defer
	defer c.RetData(resp)

	// data from session

	// data from mysql
	var areas []models.Area

	// query database
	o:= orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Info("read data --------------------- error ")
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	// data
	resp["data"] = areas

	beego.Info("------------query data success, resp = ", resp, "num = ", num)
}
