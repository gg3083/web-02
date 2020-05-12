package model

import (
	"fmt"
	db "web-02/database"
)

type User struct {
	Id       int    `json:"id"			gorm:"primary_key" gorm:"column:id"`
	UserName string `json:"userName"    gorm:"column:user_name"`
	Birthday string `json:"birthday"    gorm:"column:birthday"`
	Sex      bool   `json:"sex"         gorm:"column:sex"`
}

// 设置Acceptor的表名为`acceptor`
func (User) TableName() string {
	return "user"
}

func (u User) UserList() (users []User, err error) {
	err = db.SqlDB.Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, err
}

func (u User) GetUser(id int) (user User, err error) {
	err = db.SqlDB.First(&user, id).Error
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, err
}

func (u User) SaveUser(entity User) (user User, err error) {
	err = db.SqlDB.Save(&entity).Error

	if err != nil {
		fmt.Println(err)
		return entity, err
	}
	return entity, err
}
