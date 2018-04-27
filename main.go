package main

import (
	_ "ihome/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_ "ihome/models"
)

func main() {
	ignoreStaticPath()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run(":8899")
}

/*
 * functional: ignore static
 */
 func ignoreStaticPath() {
 	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)

 	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
 }

 /*
  *	function: function pointer
  */
  func TransparentStatic(ctx *context.Context) {
  	orpath := ctx.Request.URL.Path
  	beego.Debug("request url:", orpath)
  	//
  	if strings.Index(orpath, "api") >=0 {
  		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)

  }