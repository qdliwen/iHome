package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"ihome/models"
	"github.com/astaxie/beego/orm"
)

/*
 * User controller
 */
 type UserController struct {
 	beego.Controller
 }

 /*
  * User controller method
  */
  func (this *UserController)Reg() {
	  resp := make(map[string]interface{})
	  defer this.retData(resp)

	  beego.Error(string(this.Ctx.Input.RequestBody))
	  // obtain page data
	  err := json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	  if err != nil {
	  	beego.Error("pasing page json data erro = ", err)
	  	return
	  }

	  beego.Info(`resp["mobile"] = `, resp["mobile"])
	  beego.Info(`resp["password"] = `, resp["password"])
	  beego.Info(`resp["sms_code"] = `, resp["sms_code"])

	  // insert to database
	  o := orm.NewOrm()
	  user := models.User{}

	  user.Password_hash = resp["password"].(string)
	  user.Name = resp["mobile"].(string)
	  user.Mobile = resp["mobile"].(string)

	  id, err := o.Insert(&user)
	  if err != nil {
	  	resp["errno"] = models.RECODE_SERVERERR
	  	resp["errmsg"]=models.RecodeText(models.RECODE_SERVERERR)
	  	return
	  }

	  beego.Info("---------------reg succes, id = ", id)

	  resp["errno"] = 0
	  resp[models.RESPONSMSGSTR] = "register successfully"
	  this.SetSession("name", user.Name)
  }

/*
*	struct methods internal
*/
func (this *UserController)retData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}