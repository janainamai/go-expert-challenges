package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	cases := []struct {
		name    string
		message string
		want    *Error
	}{
		{
			name:    "valid error",
			message: "error message",
			want: &Error{
				Message: "error message",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NewError(c.message)
			assert.Equal(t, c.want, got)
		})
	}
}
