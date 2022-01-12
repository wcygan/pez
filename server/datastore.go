package server

import (
	"errors"
	"fmt"
)

type Datastore struct {
	backingStore map[string]string
}

func (d *Datastore) Put(key, value string) error {
	d.backingStore[key] = value
	return nil
}

func (d *Datastore) Get(key string) (string, error) {
	value, ok := d.backingStore[key]
	if ok {
		return value, nil
	} else {
		return "", errKeyDoesNotExist(key)
	}
}

func errKeyDoesNotExist(key string) error {
	return errors.New(fmt.Sprintf("key %s does not exist in the datastore", key))
}
