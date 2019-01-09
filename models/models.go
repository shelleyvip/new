package models

import(
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string

}
type Article struct {
	Id int`orm:"pk;auto"`//设为主键pk
	ArtiName string`orm:"size(20)"`//设置长度：orm:"size（）
	Atime time.Time`orm:"auto_now"`//设为自增auto
	Acount int`orm:"default(0);null"`//设置允许为空null数据库默认非空，设置null之后就可以Allow null
	Acontent string
	Aimg string
 ///设置唯一键 orm:"unique"
 ///设置浮点数精度orm:"digits(12)decimals(4)"
 ///设置时间datetime data
 ///

}

func init()  {
	//设置数据库的基本信息
	orm.RegisterDataBase("default","mysql","root:123@tcp(127.0.0.1:3306)/test5?charset=utf8")
	//模型映射 也就是创建数据库表
	//映射models数据
	orm.RegisterModel(new(User),new(Article))  //new(type)也就是对象>>>(User)
	//生成表
	orm.RunSyncdb("default",false,true)

}
