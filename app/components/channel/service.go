package channel

import (
	"fmt"
	"logmontire/common/errors"
	"logmontire/storage/mongo"
)

var (
	NewChannel = make(chan *Channel)
)

func CreateChannel(args map[string]interface{}) ([]Channel, *errors.Error) {
	channel := New(args)
	channels, err := Save(channel)
	fmt.Println("channels", channels, err)
	if err != nil {
		NewChannel <- channel
	}
	return channels, err
}

func FindChannel(args map[string]interface{}) ([]Channel, *errors.Error) {
	return Find(mongo.NewSelector(args))
}

func allChannels(_ map[string]interface{}) (interface{}, *errors.Error) {
	return Find(mongo.NewEmptySelector())
}
