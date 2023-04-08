package fzf

import (
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
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
	// ItemString returns the string of the i-th item.
	ItemString(i int) string
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
	// Apply the options.
	o := defaultSearchOption
	for _, opt := range opts {
		opt(&o)
	}

	// If case-insensitive, convert the search string to lowercase.
	if !o.caseSensitive {
		search = strings.ToLower(search)
	}

	// Create a slice to store the search results.
	result := make(Matches, 0, items.Len())
	resultMutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	// Set the number of workers and chunk size.
	numWorkers := 8
	chunkSize := (items.Len() + numWorkers - 1) / numWorkers
	chunks := make(chan int, numWorkers)

	// Launch the workers.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Perform the search for each chunk.
			for start := range chunks {
				end := start + chunkSize
				if end > items.Len() {
					end = items.Len()
				}

				// Create a slice to store local matches for each chunk.
				localMatches := make(Matches, 0)

				// Perform the search for each item.
				for index := start; index < end; index++ {
					m, ok := fuzzySearch(items.ItemString(index), search, o)
					if ok {
						m.Index = index
						localMatches = append(localMatches, m)
					}
				}

				// Add the local matches to the result (while performing mutual exclusion).
				resultMutex.Lock()
				result = append(result, localMatches...)
				resultMutex.Unlock()
			}
		}()
	}

	// Assign chunks to the workers.
	for i := 0; i < items.Len(); i += chunkSize {
		chunks <- i
	}

	//Once all chunks are assigned, close the channel.
	close(chunks)

	// Wait for all workers to complete their processing.
	wg.Wait()

	// Sort the search results and return them.
	result.sort()
	return result
}

func fuzzySearch(str, search string, option searchOption) (Match, bool) {
	item := str

	// If case-insensitive, convert the item to lowercase.
	if !option.caseSensitive {
		item = strings.ToLower(item)
	}

	// Create a slice to store the matched indexes.
	matchedIndexes := make([]int, 0, utf8.RuneCountInString(search))
	j := 0

	// Check for matching between the item's characters and the search string.
	searchRunes := []rune(search)
	for i, r := range []rune(item) {
		if j < len(searchRunes) && r == searchRunes[j] {
			matchedIndexes = append(matchedIndexes, i)
			j++
		}
	}

	// Returns Match if all characters in the search string match.
	if j == len(searchRunes) {
		return Match{Str: str, MatchedIndexes: matchedIndexes}, true
	} else {
		return Match{}, false
	}
}
