package fzf

import (
	"reflect"
)

type items struct {
	items          reflect.Value
	itemFunc       func(i int) string
	itemPrefixFunc func(i int) string
}

func newItems(rv reflect.Value, itemFunc func(i int) string, itemPrefixFunc func(i int) string) (*items, error) {
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
	if is.items.Kind() == reflect.Ptr {
		return reflect.Indirect(is.items).Len()
	} else {
		return is.items.Len()
	}
}

func (is items) HasItemPrefixFunc() bool {
	return is.itemPrefixFunc != nil
}
