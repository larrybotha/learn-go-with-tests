package main

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

const (
	ErrNotFound              = DictionaryErr("no value for this key")
	ErrWordExists            = DictionaryErr("word exists")
	ErrUpdateNonExistentWord = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	var err error

	if !ok {
		err = ErrNotFound
	}

	return value, err
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	// handle Search returning any other error
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrUpdateNonExistentWord
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
