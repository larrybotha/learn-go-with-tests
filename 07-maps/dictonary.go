package main

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	return d[key]
}
