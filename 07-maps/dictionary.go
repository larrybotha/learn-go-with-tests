package dictionary

type Dictionary map[string]string

// use a const to define errors
const (
	ErrWordNotFound   = DictionaryErr("could not find the word you are looking for")
	ErrUnableToAdd    = DictionaryErr("cannot add where word exists")
	ErrUnableToUpdate = DictionaryErr("cannot update where word doesn't exist")
	ErrUnableToDelete = DictionaryErr("cannot delete where word doesn't exist")
)

// use a type alias to describe DictionaryErr
type DictionaryErr string

// implement the Error interface for DictionaryErr
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]

	if !ok {
		return "", ErrWordNotFound
	}

	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		return ErrUnableToAdd
	case ErrWordNotFound:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrUnableToUpdate
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrUnableToDelete
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
