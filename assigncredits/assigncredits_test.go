package assigncredits_test

import (
	"creditchallenge/assigncredits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignWhenIsVerySmallValue(t *testing.T) {
	c := assigncredits.New()
	a300, a500, a700, err := c.Asssign(90)
	expectedValue := int32(0)

	assert.NotNil(t, err)

	assert.Equal(t, expectedValue, a300)
	assert.Equal(t, expectedValue, a500)
	assert.Equal(t, expectedValue, a700)
}

func TestAssignWhenIsInvalidValue(t *testing.T) {
	c := assigncredits.New()
	a300, a500, a700, err := c.Asssign(310)
	expectedValue := int32(0)

	assert.NotNil(t, err)

	assert.Equal(t, expectedValue, a300)
	assert.Equal(t, expectedValue, a500)
	assert.Equal(t, expectedValue, a700)
}

func TestCalculateEachGroup(t *testing.T) {
	tests := []struct {
		Investment int32
		HasError   bool
	}{
		{
			Investment: 3000,
		},
		{
			Investment: 6700,
		},
		{
			Investment: 400,
			HasError:   true,
		},
	}

	c := assigncredits.New()
	for _, test := range tests {
		a300, a500, a700, err := c.Asssign(test.Investment)

		if test.HasError {
			assert.True(t, err != nil)
			continue
		}

		assert.True(t, err == nil)
		values := (a300 * 300) + (a500 * 500) + (a700 * 700)

		assert.Equal(t, test.Investment, values)
	}
}
