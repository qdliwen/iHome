package controllers

import (
	"github.com/astaxie/beego"
	"ihome/models"
)

/*
 * struct session
 */
type SessionController struct {
	beego.Controller
}

/*
 *	struct methods internal
 */
func (this *SessionController)retData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}

/*
 * struct methods deal routers.GetSessionData
 */
func (this *SessionController)GetSessionData()  {
	resp := make(map[string]interface{})
	defer this.retData(resp)

	user := models.User{}

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	beego.Error("name in sesssion = ", name)
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"]=user
	}
}

/*
 *	struct methods deal routers.deleteSessionData
 */
 func (this *SessionController)DeleteSessionData() {
	 resp := make(map[string]interface{})
	 defer this.retData(resp)

	 this.DelSession("name")
	 resp["errno"] = models.RECODE_OK
	 resp["errmsg"] = models.RecodeText(models.RECODE_OK)
 }