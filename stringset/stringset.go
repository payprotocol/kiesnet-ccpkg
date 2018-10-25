// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package stringset

// A Set represents an unordered set of strings
type Set map[string]bool

// New creates and returns a reference to an set
func New(items ...string) *Set {
	set := make(Set)
	for _, item := range items {
		set[item] = true
	}
	return &set
}

// AppendSlice adds strings from a slice to the set
func (s *Set) AppendSlice(items []string) {
	for _, item := range items {
		(*s)[item] = true
	}
}

// Add adds a string to the set
func (s *Set) Add(item string) {
	(*s)[item] = true
}

// Remove removes a string from the set
func (s *Set) Remove(item string) {
	delete(*s, item)
}

// Strings returns the string array
func (s *Set) Strings() []string {
	items := make([]string, 0, len(*s))
	for k := range *s {
		items = append(items, k)
	}
	return items
}

// Size returns the number of elements of the set
func (s *Set) Size() int {
	return len(*s)
}
