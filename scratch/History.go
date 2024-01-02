package main

import "github.com/google/uuid"

type History map[string]int

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (h History) Search(word string) (int, error) {
	definition, ok := h[word]
	if !ok {
		return 0, ErrNotFound
	}

	return definition, nil
}

func (h History) saveResult(word string, result int) error {
	_, err := h.Search(word)

	switch err {
	case ErrNotFound:
		h[word] = result
	case nil:
		newWord := word + "_" + uuid.NewString()
		h[newWord] = result
	default:
		return err
	}
	return nil
}
