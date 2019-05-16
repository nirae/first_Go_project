package addition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert.Equal(t, 10, Plus(5, 5), "Ca doit etre egal")
}
