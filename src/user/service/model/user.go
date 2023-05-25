package model

import (
	"gorm.io/gorm"
	"time"
)

// MigrateUserTable Migrate user table
func MigrateUserTable(db *gorm.DB) {

	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic("migrate Failed.[ERROR]=>create user table failed.")
		}
		db.Exec("ALTER TABLE `user` COMMENT 'user table'")
	}

}

type User struct {
	//primaryKey
	UserID        uint64    `gorm:"column:user_id;primaryKey;autoIncrement;comment:用户id" json:"user_id"`
	UserNo        string    `gorm:"column:user_no;not null;size:128;comment:用户编号" json:"user_no"`
	UserName      string    `gorm:"column:user_name;not null;size:190;default:'';comment:用户名称" json:"user_name"`
	Avatar        string    `gorm:"column:avatar;not null;size:500;default:'';comment:头像" json:"avatar"`
	Password      string    `gorm:"column:password;not null;size:32;default:'';comment:密码" json:"-"`
	PhoneNumber   string    `gorm:"column:phone_number;not null;size:12;comment:手机号" json:"phone_number"`
	Email         string    `gorm:"column:email;not null;size:128;default:'';comment:邮箱" json:"email"`
	UnionID       string    `gorm:"column:unionid;not null;size:190;default:'';comment:微信UnionID" json:"unionid"`
	OpenID        string    `gorm:"column:openid;not null;size:190;default:'';comment:微信公众平台openid" json:"openid"`
	MiniappOpenID string    `gorm:"column:miniapp_openid;not null;size:190;default:'';comment:小程序openid" json:"miniapp_openid"`
	AddTime       time.Time `gorm:"column:add_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"add_time"`
	UpdateTime    time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	IsDelete      uint8     `gorm:"column:is_delete;not null;default:0;comment:1-已删除" json:"is_delete"`
}

// TableName sets the table name for the User model.
func (u *User) TableName() string {
	return "user"
}
