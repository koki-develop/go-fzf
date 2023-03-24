package fzf

type Items interface {
	ItemString(int) string
	Len() int
}

type items struct {
	items Items
}

func newItems(is Items) items {
	return items{is}
}

func (is items) String(i int) string {
	return stringLinesToSpace(is.items.ItemString(i))
}

func (is items) Len() int {
	return is.items.Len()
}
