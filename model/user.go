package model

type User struct {
	Id       int
	Username string
	Age      int
	Sex      int
}

func (User) TableName() string {
	return "users"
}
