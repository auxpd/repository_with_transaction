package model

type Book struct {
	ID     int64  `gorm:"primary_key"`
	Name   string // 书名
	Author string // 作者
	Price  int64  // 价格
}

func (*Book) TableName() string {
	return "book"
}
