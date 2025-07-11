package domain

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户ID，自增长主键" json:"id"`                // 用户ID，自增长主键
	Username  string    `gorm:"column:username;not null;comment:用户名，唯一" json:"username"`                             // 用户名，唯一
	Phone     string    `gorm:"column:phone;not null;comment:手机号，唯一" json:"phone"`                                   // 手机号，唯一
	Email     string    `gorm:"column:email;not null;comment:用户邮箱，唯一" json:"email"`                                  // 用户邮箱，唯一
	Password  string    `gorm:"column:password;not null;comment:密码哈希" json:"password"`                               // 密码哈希
	Nickname  string    `gorm:"column:nickname;comment:昵称" json:"nickname"`                                          // 昵称
	AvatarURL string    `gorm:"column:avatar_url;comment:头像URL" json:"avatar_url"`                                   // 头像URL
	Bio       string    `gorm:"column:bio;comment:个人简介" json:"bio"`                                                  // 个人简介
	IsAdmin   bool      `gorm:"column:is_admin;not null;comment:是否管理员，默认false" json:"is_admin"`                      // 是否管理员，默认false
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (user *UserRepository) GetUserById(id string) (*User, error) {
	userDb := user.db
	var userDo User
	err := userDb.First(&userDo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &userDo, nil
}
