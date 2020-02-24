package datasource

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
)

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine()  *xorm.Engine {

	//数据库引擎
	engine ,err := xorm.NewEngine("mysql","root:123456@/syCms?charset=utf8")

	//根据实体创建表
	//err = engine.CreateTables(new(model.Admin))

	//同步数据库结构：主要负责对数据库实体同步更新到数据库表
	/**
	 * 自动检测和创建表，这个检测是根据表的名字
	 * 自动检测和新增中的字段，这个检测是根据字段名，同时都表中多的字段给出警告信息
	 * 自动检测，创建和删除
	 *
	 */
	//Sync2是Sync的基础上的优化的方法
	err = engine.Sync2(new(
		model.Permission),
			new(model.City),
			new(model.Admin),
			new(model.AdminPermission),
			new(model.User),
			new(model.UserOrder))
	if err != nil {
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine

}
