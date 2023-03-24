package fzf

func intContains(is []int, i int) bool {
	for _, l := range is {
		if l == i {
			return true
		}
	}
	return false
}

func intFilter(is []int, f func(i int) bool) []int {
	var rtn []int
	for _, i := range is {
		if f(i) {
			rtn = append(rtn, i)
		}
	}
	return rtn
}
