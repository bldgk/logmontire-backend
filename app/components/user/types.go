package user

import (
	"logmontire/crypto"
	"strconv"
)

type User struct {
	Username string `bson:"username"json:"username"`
	Id       string `bson:"_id"json:"_id"`
}

func (u *User) New(username string) {
	u.Id = strconv.Itoa(crypto.GenerateId())
	u.Username = username
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) SetUserName(username string) {
	u.Username = username
}

func (u *User) GetInstance() interface{} {
	return u
}
