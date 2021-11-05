package channel

import (
	"logmontire/crypto"
	"strconv"
)

type Channel struct {
	Id            string `bson:"_id"json:"_id"`
	Name          string `bson:"name"json:"name"`
	EnvironmentId string `bson:"environment_id"json:"environment_id"`
}

func (c *Channel) New(name string, envId string) {
	c.Id = strconv.Itoa(crypto.GenerateId())
	c.Name = name
	c.EnvironmentId = envId
}

func (c *Channel) GetName() string {
	return c.Name
}

func (c *Channel) GetId() string {
	return c.Id
}

func (c *Channel) GetInstance() interface{} {
	return c
}
