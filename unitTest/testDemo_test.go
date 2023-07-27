package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Jack"
	// 比较是否符合预期
	assert.Equal(t, expectOutput, output)
}

func TestRatingScore(t *testing.T) {
	assert.Equal(t, "A", RatingScore(90))
	assert.Equal(t, "B", RatingScore(80))
	assert.Equal(t, "C", RatingScore(70))
	assert.Equal(t, "D", RatingScore(60))
	assert.Equal(t, "E", RatingScore(59))
}
