package model

import (
	"gorm.io/gorm"
	"time"
)

// MigrateUserTable Migrate user table
func MigrateUserTable(db *gorm.DB) {

	m := db.Migrator()
	if !m.HasTable(&Order{}) {
		if err := m.CreateTable(&Order{}); err != nil {
			panic("migrate Failed.[ERROR]=>create user table failed.")
		}
		db.Exec("ALTER TABLE `order` COMMENT 'order table'")
	}

}

type Order struct {
	OrderId    uint64    `gorm:"column:order_id;primaryKey;autoIncrement" json:"order_id"`
	OrderNo    string    `gorm:"column:order_no;type:char(50);not null;default:'';comment:'订单号'" json:"order_no"`
	UserId     uint64    `gorm:"column:user_id;not null;default:0;comment:'用户id'" json:"user_id"`
	UserName   string    `gorm:"column:user_name;type:varchar(190);not null;default:'';comment:'用户名'" json:"user_name"`
	Remark     string    `gorm:"column:remark;type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	PriceTotal float64   `gorm:"column:price_total;type:decimal(10,2) unsigned;not null;default:0.00;comment:'订单总金额'" json:"price_total"`
	PricePay   float64   `gorm:"column:price_pay;type:decimal(10,2);not null;default:0.00;comment:'订单应付金额'" json:"price_pay"`
	PricePaid  float64   `gorm:"column:price_paid;type:decimal(10,2);not null;default:0.00;comment:'订单已付金额'" json:"price_paid"`
	Status     uint64    `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:'状态：0-未支付，1-已支付'" json:"status"`
	IsInvoice  uint64    `gorm:"column:is_invoice;type:tinyint unsigned;not null;default:0;comment:'是否已开发票.0未操作，1已申请，2申请通过'" json:"is_invoice"`
	PayTime    time.Time `gorm:"column:pay_time;type:datetime;comment:'支付时间'" json:"pay_time"`
	AddTime    time.Time `gorm:"column:add_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time"`
}

func (o Order) TableName() string {
	return "order"
}
