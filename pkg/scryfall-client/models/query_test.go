package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPointer(s string) *string {
	return &s
}

func TestQueryFormat(t *testing.T) {
	t.Run("Format search with Name", func(t *testing.T) {
		name := "Eriette"
		query := Query{
			Color: []QueryField{
				{
					Text:    "wb",
					Operand: stringPointer(">="),
				},
			},
			ColorIdentity: []QueryField{
				{
					Text: "Orzhov",
				},
			},
		}

		assert.Equal(t, "Eriette%2Bc%3E%3Dwb%2Bid%3AOrzhov", query.Format(&name))
	})
	t.Run("Format search without Name", func(t *testing.T) {
		query := Query{
			Color: []QueryField{
				{
					Text:    "wb",
					Operand: stringPointer(">="),
				},
			},
			ColorIdentity: []QueryField{
				{
					Text: "Orzhov",
				},
			},
		}

		assert.Equal(t, "c%3E%3Dwb%2Bid%3AOrzhov", query.Format(nil))
	})
}
