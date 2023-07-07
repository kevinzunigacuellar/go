package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorNotFound         = errors.New("could not find the word you were looking for")
	ErrorWordExists       = errors.New("cannot add word because it already exists")
	ErrorWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err

	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
