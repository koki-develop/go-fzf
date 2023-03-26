package fzf

import (
	"sort"
	"strings"
	"sync"
)

var (
	defaultSearchOption = searchOption{
		caseSensitive: false,
	}
)

type searchOption struct {
	caseSensitive bool
}

// Items represents a list of items to be searched.
type Items interface {
	// String returns the string of the i-th item.
	String(i int) string
	// Len returns the length of items.
	Len() int
}

// Match represents a matched string.
type Match struct {
	Str            string
	Index          int
	MatchedIndexes []int
}

// Matches is a slice of Match.
type Matches []Match

func (m Matches) sort() {
	sort.Slice(m, func(i, j int) bool {
		mi, mj := m[i].MatchedIndexes, m[j].MatchedIndexes
		li, lj := len(mi), len(mj)

		if li != lj {
			return li < lj
		}

		for k := 0; k < li; k++ {
			if mi[k] != mj[k] {
				return mi[k] < mj[k]
			}
		}

		return m[i].Index < m[j].Index
	})
}

// SearchOption represents a option for a search.
type SearchOption func(o *searchOption)

func WithSearchCaseSensitive(c bool) SearchOption {
	return func(o *searchOption) {
		o.caseSensitive = c
	}
}

// Search performs a fuzzy search for items.
func Search(items Items, search string, opts ...SearchOption) Matches {
	o := defaultSearchOption
	for _, opt := range opts {
		opt(&o)
	}

	if !o.caseSensitive {
		search = strings.ToLower(search)
	}

	result := make(Matches, 0, items.Len())
	resultMutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	numWorkers := 8
	chunkSize := (items.Len() + numWorkers - 1) / numWorkers
	chunks := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for start := range chunks {
				end := start + chunkSize
				if end > items.Len() {
					end = items.Len()
				}

				localMatches := make(Matches, 0)

				for index := start; index < end; index++ {
					item := items.String(index)

					if !o.caseSensitive {
						item = strings.ToLower(item)
					}

					matchedIndexes := make([]int, 0, len(search))
					j := 0

					for i, r := range item {
						if j < len(search) && r == rune(search[j]) {
							matchedIndexes = append(matchedIndexes, i)
							j++
						}
					}

					if j == len(search) {
						m := Match{Str: items.String(index), Index: index, MatchedIndexes: matchedIndexes}
						localMatches = append(localMatches, m)
					}
				}

				resultMutex.Lock()
				result = append(result, localMatches...)
				resultMutex.Unlock()
			}
		}()
	}

	for i := 0; i < items.Len(); i += chunkSize {
		chunks <- i
	}
	close(chunks)
	wg.Wait()

	result.sort()
	return result
}
