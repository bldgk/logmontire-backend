package mongo

import (
	"fmt"
	"logmontire/common/errors"
)

type Selector struct {
	selector map[string]interface{}
}

func (s *Selector) HasId() (bool, *errors.Error) {
	if s.selector == nil {
		return false, nil
	}

	id, exists := s.selector["_id"]
	if exists {
		if id != nil {
			return exists, nil
		} else {
			return false, EmptyObjectId()
		}
	}

	return false, nil
}

func (s *Selector) GetId() string {
	return fmt.Sprint(s.selector["_id"])
}

func (s *Selector) GetSelect() map[string]interface{} {
	return s.selector
}

func (s *Selector) GetString() string {
	for k, v := range s.selector {
		return fmt.Sprintf("%s=%s", k, v)
	}

	return ""
}

func (s *Selector) SetSelector(args map[string]interface{}) {
	s.selector = args
}

func NewSelector(args map[string]interface{}) *Selector {
	return &Selector{
		selector: args,
	}
}

func NewEmptySelector() *Selector {
	return &Selector{}
}

func NewSelectorById(id string) *Selector {
	return &Selector{
		selector: map[string]interface{}{
			"_id": id,
		},
	}
}
