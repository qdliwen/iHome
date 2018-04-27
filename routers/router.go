package routers

import (
	"ihome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})


    /*
     *router for url /api/v1.0/.../.../...
     */
    //	router for address request
    beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
    //	router for house index
    beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	//	router for sesssion
	beego.Router("api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DeleteSessionData")
	// router for user register
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")

}

