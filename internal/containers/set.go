package containers

import (
	"database/sql/driver"
	"fmt"
	"github.com/lib/pq"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type empty struct{}
type Set[T constraints.Ordered] map[T]empty
type StringSet struct {
	Set[string]
}

func NewSet[T constraints.Ordered](elements ...T) Set[T] {
	set := Set[T]{}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

func (s StringSet) Value() (driver.Value, error) {
	var a pq.StringArray
	a = s.Items()
	return a.Value()
}

func (s *StringSet) Scan(src interface{}) error {
	var a pq.StringArray
	err := a.Scan(src)
	if err != nil {
		return err
	}
	*s = NewStringSet(a...)
	return nil
}

func NewStringSet(elements ...string) StringSet {
	set := NewSet[string](elements...)
	return StringSet{
		set,
	}
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = empty{}
	}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	slices.Sort(items)
	return items
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Clear() {
	for item := range s {
		delete(s, item)
	}
}

func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

func (s Set[T]) Intersects(other Set[T]) bool {
	for item := range s {
		if other.Contains(item) {
			return true
		}
	}
	return false
}

func (s Set[T]) SubsetOf(other Set[T]) bool {
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

func (s Set[T]) SupersetOf(other Set[T]) bool {
	return other.SubsetOf(s)
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := NewSet[T]()
	for item := range s {
		if !other.Contains(item) {
			diff.Add(item)
		}
	}
	return diff
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet[T]()
	for item := range s {
		union.Add(item)
	}
	for item := range other {
		union.Add(item)
	}
	return union
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for item := range s {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	diff := s.Difference(other)
	return diff.Union(other.Difference(s))
}

func (s Set[T]) String() string {
	sb := &strings.Builder{}
	sb.WriteString("{")
	items := s.Items()
	var i = 0
	for _, item := range items {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", item))
		i++
	}
	sb.WriteString("}")
	return sb.String()
}

func (s Set[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		for item := range s {
			ch <- item
		}
		close(ch)
	}()
	return ch
}

func (s Set[T]) IteratorWithLimit(limit int) <-chan T {
	ch := make(chan T)
	go func() {
		for item := range s {
			ch <- item
			if limit--; limit == 0 {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func (s Set[T]) Clone() Set[T] {
	var out = NewSet[T]()
	for item := range s {
		out.Add(item)
	}
	return out
}
