// Package set provides methods on a set of strings.
package set

// Set is implemented as a map to empty struct.
// This is because an empty struct takes no memory.
type Set map[string]struct{}

// Add adds element e to set s.
func (s Set) Add(e string) {
	s[e] = struct{}{}
}

// In checks if element e is in set s.
func (s Set) In(e string) bool {
	_, ok := s[e]
	return ok
}

// And returns the intersection between set t and set s.
// The intersection is a new set.
func (s Set) And(t Set) Set {
	u := Set{}
	for se := range s {
		if t.In(se) {
			u.Add(se)
		}
	}
	return u
}

// Or returns the union of set t and set s.
func (s Set) Or(t Set) Set {
	u := Set{}
	for se := range s {
		u.Add(se)
	}
	for te := range t {
		u.Add(te)
	}
	return u
}

// Delete deletes element e from set s.
func (s Set) Delete(e string) Set {
	delete(s, e)
	return s
}

// Subset checks if t is a proper subset of s.
func (s Set) Subset(t Set) bool {
	if len(t) >= len(s) {
		return false
	}
	if t.Equal(s.And(t)) {
		return true
	}
	return false
}

// Equal checks if set t is equal to set s.
func (s Set) Equal(t Set) bool {
	return len(s.And(t)) == len(s.Or(t))
}
