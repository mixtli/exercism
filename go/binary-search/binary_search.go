package binarysearch

import "fmt"

const testVersion = 1

// SearchInts returns the index of needle in haystack
func SearchInts(haystack []int, needle int) int {
	first := 0
	last := len(haystack) - 1
	for first <= last {
		middle := (first + last) / 2
		if haystack[middle] < needle {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	return first
}

// Message returns an english message describing the index position
func Message(haystack []int, needle int) string {
	index := SearchInts(haystack, needle)
	atStart, atEnd := index == 0, index == len(haystack)-1
	beyondEnd := index == len(haystack)
	found := !beyondEnd && haystack[index] == needle

	switch {
	case 0 == len(haystack):
		return "slice has no values"
	case found && atStart:
		return fmt.Sprintf("%d found at beginning of slice", needle)
	case found && atEnd:
		return fmt.Sprintf("%d found at end of slice", needle)
	case found:
		return fmt.Sprintf("%d found at index %d", needle, index)
	case !found && atStart:
		return fmt.Sprintf("%d < all values", needle)
	case !found && beyondEnd:
		return fmt.Sprintf("%d > all %d values", needle, len(haystack))
	case !found:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
			needle, haystack[index-1], index-1, haystack[index], index)
	default:
		return ""
	}
}
