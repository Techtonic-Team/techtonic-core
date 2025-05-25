package subduction_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/techtonic-team/tdk/techtonic-core/pkg/errors/subduction"
)

func TestError(t *testing.T) {
	err := fmt.Errorf("file not found")
	wrapped := subduction.Wrap(err, "config load failed")

	assert.EqualError(t, wrapped, "config load failed: file not found")
	assert.True(t, errors.Is(wrapped, err))
}
