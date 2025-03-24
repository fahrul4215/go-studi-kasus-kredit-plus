package pagination

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLimit(t *testing.T) {
	tests := []struct {
		tag                  string
		page, perPage, total int
		expectedLimit        int
	}{
		{"t1", 1, 0, 50, DefaultLimit},
		{"t2", 1, -1, 50, DefaultLimit},
		{"t3", 1, DefaultLimit, 50, DefaultLimit},
	}

	for _, test := range tests {
		p := New(test.page, test.perPage, test.total, "", "", "")
		assert.Equal(t, test.expectedLimit, p.Limit, test.tag)
	}
}

func TestNewFromRequest_DefaultLimit(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com?page=2", bytes.NewBufferString(""))
	p := NewFromRequest(req)
	assert.Equal(t, 2, p.Page)
	assert.Equal(t, DefaultLimit, p.Limit)
}
