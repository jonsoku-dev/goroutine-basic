package myDict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errorNotFound = errors.New("Not found word")
var errorWordExists = errors.New("already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errorNotFound {
		d[word] = def
	} else if err == nil {
		return errorWordExists
	}
	return nil
}
