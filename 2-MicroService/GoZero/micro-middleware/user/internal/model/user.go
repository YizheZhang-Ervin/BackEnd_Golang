package model

type User struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	Gender string `db:"gender"`
}

func (User) TableName() string {
	return "user"
}
