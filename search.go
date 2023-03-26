package fzf

import (
	"sort"
	"sync"
)

type match struct {
	Str            string
	Index          int
	MatchedIndexes []int
}

type matches []match

func (m matches) Sort() {
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

func fuzzySearch(items *items, search string) matches {
	result := make(matches, 0, items.Len())
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

				localMatches := make(matches, 0)

				for index := start; index < end; index++ {
					item := items.String(index)
					matchedIndexes := make([]int, 0, len(search))
					j := 0

					for i, r := range item {
						if j < len(search) && r == rune(search[j]) {
							matchedIndexes = append(matchedIndexes, i)
							j++
						}
					}

					if j == len(search) {
						m := match{Str: item, Index: index, MatchedIndexes: matchedIndexes}
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

	result.Sort()
	return result
}
