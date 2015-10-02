package iterator

import (
	"errors"
	"fmt"
	"reflect"
)

type Iterator interface {
	Reset()
	Next() interface{}
	HasNext() bool
}

type iterator struct {
	index int
	slice reflect.Value
}

func New(slice interface{}) Iterator {
	value := reflect.ValueOf(slice)

	if value.Kind() != reflect.Slice || value.Kind() == reflect.Ptr && value.Elem().Kind() != reflect.Slice {
		panic(errors.New(fmt.Sprintf("Expeted a slice or a pointer to a slice, got %v", value.Kind())))
	}

	return &iterator{
		slice: value,
	}
}

func (i *iterator) Reset() {
	i.index = 0
}

func (i *iterator) HasNext() bool {
	return i.index < i.slice.Len()
}

func (i *iterator) Next() interface{} {
	if i.HasNext() {
		item := i.slice.Index(i.index).Interface()
		i.index++
		return item
	}
	return nil
}
