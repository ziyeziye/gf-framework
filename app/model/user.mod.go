package model

import (
	"database/sql"
)

var UserModel = Model(User{})

type User struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	UserNick   string `json:"user_nick"`
	UserPass   string `json:"user_pass"`
	UserImg    string `json:"user_img"`
	LoginTime  string `json:"login_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (u User) TableName() string {
	return GetPrefix() + "users"
}

func GetUser(id int) (user *User, err error) {
	err = UserModel.Where("id", id).Struct(user)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetUsers(maps map[string]interface{}) (users []*User, err error) {
	query, err := ModelSearch(User{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(users)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetUserTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = UserModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteUser(user *User) (err error) {
	_, err = UserModel.Where(user).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateUser(user *User, data map[string]interface{}) (err error) {
	_, err = UserModel.Where(user).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddUser(user *User) (err error) {
	_, err = UserModel.Data(user).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
