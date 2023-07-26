package unittest

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestProcessFirstLine(t *testing.T) {
	input := ProcessFirstLine()
	assert.Equal(t, "line00", input)
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(readFirstLine, func() string {
		return "Line111"
	})
	defer monkey.Unpatch(readFirstLine)
	assert.Equal(t, "Line000", ProcessFirstLine())
}
