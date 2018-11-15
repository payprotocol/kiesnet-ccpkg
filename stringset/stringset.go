// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package stringset

import "encoding/json"

// A Set represents an unordered set of strings
type Set map[string]bool

// New creates and returns a reference to an set
func New(items ...string) Set {
	set := make(Set)
	for _, item := range items {
		set[item] = true
	}
	return set
}

// Add adds a string to the set
func (s Set) Add(item string) {
	if s != nil {
		s[item] = true
	}
}

// AppendSet adds strings of the 'set' to the set
func (s Set) AppendSet(set Set) {
	if s != nil {
		for k := range set {
			s[k] = true
		}
	}
}

// AppendSlice adds strings from a slice to the set
func (s Set) AppendSlice(items []string) {
	if s != nil {
		for _, item := range items {
			s[item] = true
		}
	}
}

// Contains reports whether item is within the set
func (s Set) Contains(item string) bool {
	return s[item]
}

// Remove removes a string from the set
func (s Set) Remove(item string) {
	delete(s, item)
}

// Size returns the number of elements of the set
func (s Set) Size() int {
	return len(s)
}

// Strings returns the string array
func (s Set) Strings() []string {
	items := make([]string, 0, len(s))
	for k := range s {
		items = append(items, k)
	}
	return items
}

// MarshalJSON implements the json.Marshaler interface
func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Strings())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s Set) UnmarshalJSON(text []byte) error {
	var items []string
	if err := json.Unmarshal(text, &items); err != nil {
		return err
	}
	s.AppendSlice(items)
	return nil
}
