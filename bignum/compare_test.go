package bignum

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsNotNegative(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{
			give: "-0.0000001",
			want: false,
		},
		{
			give: "0",
			want: true,
		},
		{
			give: "99999999999999999999",
			want: true,
		},
		{
			give: "ddddddddd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			result := IsNotNegative(tt.give)
			assert.Equal(t, result, tt.want, "")
		})
	}
}

func TestIsNegative(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{
			give: "-0.0000001",
			want: true,
		},
		{
			give: "0",
			want: false,
		},
		{
			give: "99999999999999999999",
			want: false,
		},
		{
			give: "ddddddddd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			result := IsNegative(tt.give)
			assert.Equal(t, result, tt.want, "")
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{
			give: "-0.0000001",
			want: false,
		},
		{
			give: "0",
			want: false,
		},
		{
			give: "99999999999999999999",
			want: true,
		},
		{
			give: "ddddddddd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			result := IsPositive(tt.give)
			assert.Equal(t, result, tt.want, "")
		})
	}
}

func TestIsNotPositive(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{
			give: "-0.0000001",
			want: true,
		},
		{
			give: "0",
			want: true,
		},
		{
			give: "99999999999999999999",
			want: false,
		},
		{
			give: "ddddddddd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			result := IsNotPositive(tt.give)
			assert.Equal(t, result, tt.want, "")
		})
	}
}
