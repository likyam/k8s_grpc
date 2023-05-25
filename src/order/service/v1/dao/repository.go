package dao

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"likyam.cn/src/order/service/model"
)

// Repository Repository
type Repository struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewRepository New Repository
func NewRepository(db *gorm.DB, redis *redis.Client) *Repository {
	return &Repository{
		db:    db,
		redis: redis,
	}
}

// UserModel User model
func (r *Repository) OrderModel() *gorm.DB {
	return r.db.Table("order")
}

// Info 获取用户信息
func (r *Repository) Info(orderId uint64) (*model.Order, error) {

	order := &model.Order{}
	result := r.OrderModel().Where("order_id = ?", orderId).First(&order)
	if result.Error != nil {
		return nil, errors.New("查询订单数据失败，错误：" + result.Error.Error())
	}
	return order, nil

}
