package dao

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"likyam.cn/src/user/service/model"
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
func (r *Repository) UserModel() *gorm.DB {
	return r.db.Table("user")
}

// Info 获取用户信息
func (r *Repository) Info(userId uint64) (*model.User, error) {

	user := &model.User{}
	result := r.UserModel().Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, errors.New("查询用户数据失败，错误：" + result.Error.Error())
	}
	return user, nil

}
