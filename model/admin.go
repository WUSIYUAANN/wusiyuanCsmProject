package model

import "time"

//定义管理员结构体
type Admin struct {
	//如果filed名称为Id,类型为int64，并没有定义tag,则会被xrom视为主键
	AdminId      int64     `xorm:"pk autoincr" json:"id"`    //主键 自增
	AdminName    string    `xorm:"varchar(32)" json:"admin_name"`
	CreatTime    time.Time `xorm:"DataTime" json:"creat_time"`
	Status       int64     `xorm:"default 0" json:"status"`
	Avatar       string    `xorm:"varchar(255)" json:"avatar"`
	Pwd          string    `xorm:"varchar(255)" json:"pwd"`      //管理员密码
	CityName     string    `xorm:"varchar(12)" json:"city_name"` //管理员所在的城市名称
	City         *City     `xorm:"- <- ->"`   //所对应的城市结构体
}

/**
 * 从Admin数据库实体转换为前端请求的json数据格式
 */

func (this *Admin) AdminToRespDesc() interface{}  {
	respDesc := map[string]interface{}{
		"user_name":    this.AdminName,
		"id":			this.AdminId,
		"create_time":  this.CreatTime,
		"status":		this.Status,
		"avatar":		this.Avatar,

	}
}