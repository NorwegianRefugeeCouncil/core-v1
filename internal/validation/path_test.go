package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Path
		want string
	}{
		{
			name: "nil",
			p:    nil,
			want: "<nil>",
		}, {
			name: "root",
			p:    &Path{name: "root"},
			want: "root",
		}, {
			name: "with child",
			p:    NewPath("root", "child"),
			want: "root.child",
		}, {
			name: "with index",
			p:    NewPath("root").Index(1),
			want: "root[1]",
		}, {
			name: "with key",
			p:    NewPath("root").Key("key"),
			want: "root[key]",
		}, {
			name: "with index and key",
			p:    NewPath("root").Index(1).Key("key"),
			want: "root[1][key]",
		}, {
			name: "with child and index",
			p:    NewPath("root", "child").Index(1),
			want: "root.child[1]",
		}, {
			name: "with child and key",
			p:    NewPath("root", "child").Key("key"),
			want: "root.child[key]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Path.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Root(t *testing.T) {
	tests := []struct {
		name string
		p    *Path
		want *Path
	}{
		{
			name: "nil",
			p:    nil,
			want: nil,
		}, {
			name: "root",
			p:    &Path{name: "root"},
			want: &Path{name: "root"},
		}, {
			name: "with child",
			p:    NewPath("root", "child"),
			want: &Path{name: "root"},
		}, {
			name: "with index",
			p:    NewPath("root").Index(1),
			want: &Path{name: "root"},
		}, {
			name: "with key",
			p:    NewPath("root").Key("key"),
			want: &Path{name: "root"},
		}, {
			name: "with index and key",
			p:    NewPath("root").Index(1).Key("key"),
			want: &Path{name: "root"},
		}, {
			name: "with child and index",
			p:    NewPath("root", "child").Index(1),
			want: &Path{name: "root"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.Root())
		})
	}
}
