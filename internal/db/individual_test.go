package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTableSchemaQuery(t *testing.T) {
	t.Run("Get Table schema query", func(t *testing.T) {
		query := buildTableSchemaQuery()
		assert.Equal(t, query, "SELECT column_name as Name, udt_name as SQLType, column_default as Default FROM information_schema.columns WHERE table_name = 'individual_registrations';")
	})
}

