package utils

import "github.com/thefuga/go-collections"

func [T any]shuffle(s []T) []T {
	return collections.Shuffle(s)
}
