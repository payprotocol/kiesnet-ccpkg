// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package stringset

import "encoding/json"

// A Set represents an unordered set of strings
type Set struct {
	m map[string]bool
}

// New creates and returns a reference to an set
func New(items ...string) *Set {
	set := &Set{make(map[string]bool, len(items))}
	set.AppendSlice(items)
	return set
}

func (s *Set) safeMap() {
	if nil == s.m {
		s.m = make(map[string]bool)
	}
}

// Add adds a string to the set
func (s *Set) Add(item string) {
	s.safeMap()
	s.m[item] = true
}

// AppendSet adds strings of the 'set' to the set
func (s *Set) AppendSet(set *Set) {
	s.safeMap()
	for k := range s.m {
		s.m[k] = true
	}
}

// AppendSlice adds strings from a slice to the set
func (s *Set) AppendSlice(items []string) {
	s.safeMap()
	for _, item := range items {
		s.m[item] = true
	}
}

// Contains reports whether item is within the set
func (s *Set) Contains(item string) bool {
	return s.m[item]
}

// Map returns strings map
func (s *Set) Map() map[string]bool {
	return s.m
}

// Remove removes a string from the set
func (s *Set) Remove(item string) {
	delete(s.m, item)
}

// Size returns the number of elements of the set
func (s *Set) Size() int {
	return len(s.m)
}

// Strings returns the string array
func (s *Set) Strings() []string {
	items := make([]string, 0, s.Size())
	for k := range s.m {
		items = append(items, k)
	}
	return items
}

// MarshalJSON implements the json.Marshaler interface
func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Strings())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *Set) UnmarshalJSON(text []byte) error {
	var items []string
	if err := json.Unmarshal(text, &items); err != nil {
		return err
	}
	s.AppendSlice(items)
	return nil
}
