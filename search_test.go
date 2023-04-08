package fzf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fuzzySearchArgs struct {
	str    string
	search string
	option searchOption
}

func assertFuzzySearch(t *testing.T, args fuzzySearchArgs, want Match, ok bool) {
	got1, got2 := fuzzySearch(args.str, args.search, args.option)

	if ok {
		assert.True(t, got2)
		assert.Equal(t, want, got1)
	} else {
		assert.False(t, got2)
	}
}

func Test_fuzzySearch(t *testing.T) {
	tests := []struct {
		str            string
		search         string
		matchedIndexes []int
	}{
		{str: "abc", search: "", matchedIndexes: []int{}},
		{str: "abc", search: "a", matchedIndexes: []int{0}},
		{str: "abc", search: "ab", matchedIndexes: []int{0, 1}},
		{str: "abc", search: "ac", matchedIndexes: []int{0, 2}},
		{str: "abc", search: "abc", matchedIndexes: []int{0, 1, 2}},
		{str: "abc", search: "b", matchedIndexes: []int{1}},
		{str: "abc", search: "bc", matchedIndexes: []int{1, 2}},
		{str: "abc", search: "c", matchedIndexes: []int{2}},
		{str: "abc", search: "cba"},
		{str: "abc", search: "d"},
		{str: "abc", search: "abcd"},

		{str: "こんにちは", search: "", matchedIndexes: []int{}},
		{str: "こんにちは", search: "こ", matchedIndexes: []int{0}},
		{str: "こんにちは", search: "こん", matchedIndexes: []int{0, 1}},
		{str: "こんにちは", search: "こんに", matchedIndexes: []int{0, 1, 2}},
		{str: "こんにちは", search: "こんにち", matchedIndexes: []int{0, 1, 2, 3}},
		{str: "こんにちは", search: "こんにちは", matchedIndexes: []int{0, 1, 2, 3, 4}},
		{str: "こんにちは", search: "ん", matchedIndexes: []int{1}},
		{str: "こんにちは", search: "んに", matchedIndexes: []int{1, 2}},
		{str: "こんにちは", search: "んにち", matchedIndexes: []int{1, 2, 3}},
		{str: "こんにちは", search: "んにちは", matchedIndexes: []int{1, 2, 3, 4}},
		{str: "こんにちは", search: "こには", matchedIndexes: []int{0, 2, 4}},

		{str: "xaxbxc", search: "a", matchedIndexes: []int{1}},
		{str: "xaxbxc", search: "ab", matchedIndexes: []int{1, 3}},
		{str: "xaxbxc", search: "ac", matchedIndexes: []int{1, 5}},
		{str: "xaxbxc", search: "abc", matchedIndexes: []int{1, 3, 5}},
		{str: "xaxbxc", search: "b", matchedIndexes: []int{3}},
		{str: "xaxbxc", search: "bc", matchedIndexes: []int{3, 5}},
		{str: "xaxbxc", search: "c", matchedIndexes: []int{5}},
		{str: "xaxbxc", search: "cba"},
		{str: "xaxbxc", search: "d"},
		{str: "xaxbxc", search: "abcd"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assertFuzzySearch(
				t,
				fuzzySearchArgs{str: tt.str, search: tt.search},
				Match{Str: tt.str, MatchedIndexes: tt.matchedIndexes},
				tt.matchedIndexes != nil,
			)
		})
	}
}

func Test_fuzzySearch_caseSensitive(t *testing.T) {
	tests := []struct {
		str            string
		search         string
		matchedIndexes []int
	}{
		{str: "abc", search: "abc", matchedIndexes: []int{0, 1, 2}},
		{str: "abc", search: "Abc"},
		{str: "abc", search: "ABC"},

		{str: "Abc", search: "abc"},
		{str: "Abc", search: "Abc", matchedIndexes: []int{0, 1, 2}},
		{str: "Abc", search: "ABC"},

		{str: "ABC", search: "abc"},
		{str: "ABC", search: "Abc"},
		{str: "ABC", search: "ABC", matchedIndexes: []int{0, 1, 2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assertFuzzySearch(
				t,
				fuzzySearchArgs{str: tt.str, search: tt.search, option: searchOption{caseSensitive: true}},
				Match{Str: tt.str, MatchedIndexes: tt.matchedIndexes},
				tt.matchedIndexes != nil,
			)
		})
	}
}
