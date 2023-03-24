package fzf

import (
	"fmt"
	"reflect"
)

type items struct {
	items          reflect.Value
	itemFunc       func(i int) string
	itemPrefixFunc func(i int) string
}

func newItems(is interface{}, itemFunc func(i int) string, itemPrefixFunc func(i int) string) (*items, error) {
	rv := reflect.ValueOf(is)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("items must be a slice, but got %T", is)
	}

	return &items{
		items:          rv,
		itemFunc:       itemFunc,
		itemPrefixFunc: itemPrefixFunc,
	}, nil
}

func (is items) String(i int) string {
	return stringLinesToSpace(is.itemFunc(i))
}

func (is items) Len() int {
	return is.items.Len()
}

func (is items) HasItemPrefixFunc() bool {
	return is.itemPrefixFunc != nil
}
