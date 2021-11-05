package channel

import (
	"logmontire/common/errors"
	"logmontire/storage"
	"logmontire/storage/mongo"
)

var (
	collection = "channels"
)

func New(args map[string]interface{}) *Channel {
	channel := &Channel{}
	channel.New(args["name"].(string), args["environment_id"].(string))
	return channel
}

func Save(channel *Channel) ([]Channel, *errors.Error) {
	var channels []Channel

	if err := storage.Save(collection, channel, &channels); err != nil {
		return nil, err
	}

	return channels, nil
}

func Find(selector *mongo.Selector) ([]Channel, *errors.Error) {
	var channels []Channel

	if err := storage.Find(collection, selector, &channels); err != nil {
		return nil, err
	}

	return channels, nil
}

func FindById(id string) (*Channel, *errors.Error) {
	var channels []Channel

	if err := storage.FindById(collection, id, &channels); err != nil {
		return nil, err
	}

	return &channels[0], nil
}
