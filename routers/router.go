package routers

import (
	"news/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HeadLogin")
    beego.Router("index",&controllers.MainController{},"get:ShowIndex")
	beego.Router("/addArticle",&controllers.MainController{},"get:ShowAdd;post:HeadAdd")
	//beego.Router("/addArtacle", &controllers.MainController{},"get:ShowAdd;post:HandleAdd")




}
