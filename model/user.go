package model

import (
	"fmt"
	db "web-02/database"
)

type User struct {
	Id       int    `json:"id"			gorm:"primary_key" gorm:"column:id"`
	UserName string `json:"userName"    gorm:"column:user_name"'`
	Birthday string `json:"birthday"    gorm:"column:birthday"`
	Sex      bool   `json:"sex"         gorm:"column:sex"`
	Session *Session `json:"-" gorm:"-"`

}

func (u User) UserList() (users []User, err error) {
	db.SqlDB.Find()
	rows, err := db.SqlDB.Query("select id,user_name,birthday,sex from test")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	fmt.Println("---------------")
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.UserName, &user.Birthday, &user.Sex)
		users = append(users, user)
	}
	return users, err
}

func (u User) GetUser(id int) (user User, err error) {
	err = db.SqlDB.QueryRow("select id,user_name,birthday,sex from test where id = ?", id).Scan(&user.Id, &user.UserName, &user.Birthday, &user.Sex)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	fmt.Println("---------------")
	return user, err
}

func (u User) UpdateUser(entity User) (user User, err error) {
	query, err := db.SqlDB.Exec("update test set user_name = ? ,birthday = ? , sex = ? where id = ?", entity.UserName, entity.Birthday, entity.Sex, entity.Id)

	if err != nil {
		fmt.Println(err)
		return entity, err
	}
	fmt.Println("---------------")
	fmt.Println(&query)
	return entity, err
}
