package fzf

// Items is a list of items to be searched by the Fuzzy Finder.
type Items interface {
	// ItemString returns the string of the i-th item in Items.
	ItemString(i int) string

	// Len returns the number of items in Items.
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
