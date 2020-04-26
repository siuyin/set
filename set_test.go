package set

import "testing"

func TestAdd(t *testing.T) {
	s := Set{}
	s.Add("a")
	if !s.In("a") {
		t.Error("a should be in set s")
	}
	if len(s) != 1 {
		t.Error("unexpected length")
	}
}

func TestAnd(t *testing.T) {
	p := Set{}
	p.Add("a")
	p.Add("b")
	q := Set{}
	q.Add("b")
	q.Add("c")
	r := p.And(q)
	if !r.In("b") {
		t.Error("expected b to be in p and q")
	}
}
func TestOrEqual(t *testing.T) {
	p := Set{}
	p.Add("a")
	q := Set{}
	q.Add("a")
	q.Add("b")
	r := p.Or(q)
	if !r.Equal(q) {
		t.Error("expected p union q = q")
	}
}

func TestDelete(t *testing.T) {
	p := Set{}
	p.Add("a")
	p.Delete("a")
	if !p.Equal(Set{}) {
		t.Error("expected p to be empty set")
	}
}

func TestSubset(t *testing.T) {
	p := Set{}
	p.Add("a")
	q := Set{}
	q.Add("a")
	q.Add("b")
	if !q.Subset(p) {
		t.Error("expected p is a subset of q")
	}
}
