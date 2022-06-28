package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	s := NewStringSet()
	assert.False(t, s.Contains("a"))
	s.Add("a")
	assert.True(t, s.Contains("a"))
}

func TestSetRemove(t *testing.T) {
	s := NewStringSet("a")
	s.Remove("a")
	assert.False(t, s.Contains("a"))
}

func TestSetContains(t *testing.T) {
	s := NewStringSet("a")
	assert.True(t, s.Contains("a"))
	assert.False(t, s.Contains("b"))
}

func TestSetLen(t *testing.T) {
	s := NewStringSet("a")
	assert.Equal(t, 1, s.Len())
}

func TestSetItems(t *testing.T) {
	s := NewStringSet("a")
	assert.ElementsMatch(t, []string{"a"}, s.Items())
}

func TestSetIsEmpty(t *testing.T) {
	s := NewStringSet()
	assert.True(t, s.IsEmpty())
	s.Add("a")
	assert.False(t, s.IsEmpty())
}

func TestSetClear(t *testing.T) {
	s := NewStringSet("a")
	assert.False(t, s.IsEmpty())
	s.Clear()
	assert.True(t, s.IsEmpty())
}

func TestSetEquals(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.True(t, s1.Equals(s2))
	s2.Add("b")
	assert.False(t, s1.Equals(s2))
}

func TestSetIntersects(t *testing.T) {
	s1 := NewStringSet()
	s2 := NewStringSet()
	assert.False(t, s1.Intersects(s2))
	s1.Add("a")
	assert.False(t, s1.Intersects(s2))
	s2.Add("a")
	assert.True(t, s1.Intersects(s2))
	s2.Add("b")
	assert.True(t, s1.Intersects(s2))
}

func TestSetSubsetOf(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.True(t, s1.SubsetOf(s2))
	s2.Add("b")
	assert.False(t, s2.SubsetOf(s1))
	assert.True(t, s1.SubsetOf(s2))
}

func TestSetSupersetOf(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.True(t, s1.SupersetOf(s2))
	s2.Add("b")
	assert.True(t, s2.SupersetOf(s1))
	assert.False(t, s1.SupersetOf(s2))
}

func TestSetDifference(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.ElementsMatch(t, []string{}, s1.Difference(s2).Items())
	s1.Add("b")
	assert.ElementsMatch(t, []string{"b"}, s1.Difference(s2).Items())
	s2.Add("b")
	assert.ElementsMatch(t, []string{}, s1.Difference(s2).Items())
}

func TestSetUnion(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.ElementsMatch(t, []string{"a"}, s1.Union(s2).Items())
	s1.Add("b")
	assert.ElementsMatch(t, []string{"a", "b"}, s1.Union(s2).Items())
	s2.Add("b")
	assert.ElementsMatch(t, []string{"a", "b"}, s1.Union(s2).Items())
}

func TestSetIntersection(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.ElementsMatch(t, []string{"a"}, s1.Intersection(s2).Items())
	s1.Add("b")
	assert.ElementsMatch(t, []string{"a"}, s1.Intersection(s2).Items())
	s2.Add("b")
	assert.ElementsMatch(t, []string{"a", "b"}, s1.Intersection(s2).Items())
}

func TestSetSymmetricDifference(t *testing.T) {
	s1 := NewStringSet("a")
	s2 := NewStringSet("a")
	assert.ElementsMatch(t, []string{}, s1.SymmetricDifference(s2).Items())
	s1.Add("b")
	assert.ElementsMatch(t, []string{"b"}, s1.SymmetricDifference(s2).Items())
	s2.Add("c")
	assert.ElementsMatch(t, []string{"b", "c"}, s1.SymmetricDifference(s2).Items())
}

func TestSetString(t *testing.T) {
	s := NewStringSet("a", "b")
	assert.Equal(t, "{a, b}", s.String())
}

func TestSetIterator(t *testing.T) {
	s := NewStringSet("a", "b")
	it := s.Iterator()
	seen := NewStringSet()
	for elem := range it {
		seen.Add(elem)
	}
	assert.ElementsMatch(t, []string{"a", "b"}, seen.Items())
}

func TestSetIteratorLimit(t *testing.T) {
	s := NewStringSet("a", "b")
	it := s.IteratorWithLimit(2)
	seen := NewStringSet()
	for elem := range it {
		seen.Add(elem)
		if seen.Len() == 2 {
			break
		}
	}
	assert.ElementsMatch(t, []string{"a", "b"}, seen.Items())
}
