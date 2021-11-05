package environment

import (
	"logmontire/crypto"
	"strconv"
)

type Environment struct {
	Name  string   `bson:"name"json:"name"`
	Token string   `bson:"token"json:"token"`
	Users []string `bson:"users"json:"users"`
	Id    string   `bson:"_id"json:"_id"`
}

func (e *Environment) New(name string, users []string) {
	e.Id = strconv.Itoa(crypto.GenerateId())
	e.Name = name
	e.Users = users
}

func (e *Environment) GenerateToken() {
	//e.Token = base64.Encode()
}

func (e *Environment) GetName() string {
	return e.Name
}

func (e *Environment) GetId() string {
	return e.Id
}

func (e *Environment) GetInstance() interface{} {
	return e
}
