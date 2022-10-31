package validation

import (
	"bytes"
	"fmt"
	"strconv"
)

// Path represents the path from some root to a particular field
type Path struct {
	// name is the name of the field or "" if this is an index or key
	name string
	// index is the index or key of the field or "" if this is a name
	index string
	// parent is the parent path or nil if this is the root
	parent *Path
}

// NewPath creates a new root Path
func NewPath(name string, moreNames ...string) *Path {
	p := &Path{name: name}
	for _, n := range moreNames {
		p = &Path{name: n, parent: p}
	}
	return p
}

// Root returns the root of the path
func (p *Path) Root() *Path {
	for ; p != nil && p.parent != nil; p = p.parent {
		// noop
	}
	return p
}

// Child returns a new path with the given name as a child of this path
func (p *Path) Child(name string, moreNames ...string) *Path {
	r := NewPath(name, moreNames...)
	r.Root().parent = p
	return r
}

// Index returns a new path with the given index as a child of this path
func (p *Path) Index(index int) *Path {
	return &Path{index: strconv.Itoa(index), parent: p}
}

// Key returns a new path with the given key as a child of this path
func (p *Path) Key(key string) *Path {
	return &Path{index: key, parent: p}
}

// String returns a string representation of the path
func (p *Path) String() string {
	if p == nil {
		return "<nil>"
	}
	elems := make([]*Path, 0)
	for ; p != nil; p = p.parent {
		elems = append(elems, p)
	}
	buf := bytes.NewBuffer(nil)
	for i := range elems {
		p := elems[len(elems)-1-i]
		if p.parent != nil && len(p.name) > 0 {
			buf.WriteByte('.')
		}
		if len(p.name) > 0 {
			buf.WriteString(p.name)
		} else {
			fmt.Fprintf(buf, "[%s]", p.index)
		}
	}
	return buf.String()
}
