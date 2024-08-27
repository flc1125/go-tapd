package tapd

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_PriorityLabel(t *testing.T) {
	tests := []struct {
		name   string
		label  PriorityLabel
		expect string
	}{
		{"", "", ""},
		{"", High, "High"},
		{"", Middle, "Middle"},
		{"", Low, "Low"},
		{"", NiceToHave, "Nice To Have"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.label.String())
		})
	}
}

func TestTypes_Order(t *testing.T) {
	tests := []struct {
		name  string
		order *Order
		want  string
	}{
		{"", NewOrder("created"), `"created asc"`},
		{"", NewOrder("created", OrderAsc), `"created asc"`},
		{"", NewOrder("created", OrderDesc), `"created desc"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.order)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(bytes))

			var o Order
			err = json.Unmarshal(bytes, &o)
			assert.NoError(t, err)
			assert.Equal(t, tt.order, &o)
		})
	}
}

func TestTypes_Order_Custom(t *testing.T) {
	type Demo struct {
		Order *Order `json:"order"`
	}

	demo := &Demo{
		Order: NewOrder("id", OrderAsc),
	}

	bytes, err := json.Marshal(demo)
	assert.NoError(t, err)
	assert.Equal(t, `{"order":"id asc"}`, string(bytes))
}
