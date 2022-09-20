package model

type User struct {
	ID       int64  `gorm:"primary_key"`
	Username string // 用户名
	Password string // 密码
	Email    string // 邮箱
	Phone    string // 手机号
	Avatar   string // 头像
}

func (*User) TableName() string {
	return "user"
}
