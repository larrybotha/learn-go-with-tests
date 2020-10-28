package main

import "errors"

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

var ErrNoValue = errors.New("no value for this key")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	var err error

	if !ok {
		err = ErrNoValue
	}

	return value, err
}
