package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBatch(t *testing.T) {
	tests := []struct {
		name      string
		batchSize int
		all       []string
		want      [][]string
		wantErr   bool
	}{
		{
			name:      "empty",
			batchSize: 10,
			all:       []string{},
			want:      [][]string{},
		}, {
			name:      "batch size 1",
			batchSize: 1,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a"}, {"b"}, {"c"}},
		}, {
			name:      "batch size 2",
			batchSize: 2,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b"}, {"c"}},
		}, {
			name:      "batch size 3",
			batchSize: 3,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b", "c"}},
		}, {
			name:      "batch size 4",
			batchSize: 4,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b", "c"}},
		}, {
			name:      "batch size 0",
			batchSize: 0,
			all:       []string{"a", "b", "c"},
			wantErr:   true,
		}, {
			name:      "batch size -1",
			batchSize: -1,
			all:       []string{"a", "b", "c"},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = make([][]string, 0)
			err := batch(tt.batchSize, tt.all, func(batch []string) error {
				got = append(got, batch)
				return nil
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
