package healthz

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWarning(t *testing.T) {
	assert.False(t, IsWarning(errors.New("foo")))
	assert.True(t, IsWarning(Warn("foo")))
	assert.True(t, IsWarning(Warnf("foo %d", 42)))
	assert.True(t, IsWarning(fmt.Errorf("wrapped: %w", Warn("foo"))))
	assert.False(t, IsWarning(fmt.Errorf("not wrapped: %v", Warn("foo"))))
	assert.True(t, IsWarning(Warnf("can wrap errors: %w", errors.New("foo"))))
}

func TestWarning_Unwrap(t *testing.T) {
	orig := errors.New("test")
	w := Warning{Err: orig}
	err := errors.Unwrap(w)
	assert.Equal(t, orig, err)
}

func TestWarning_Unwrap2(t *testing.T) {
	orig := errors.New("test")
	w := Warnf("test wrap: %w", orig)
	err := errors.Unwrap(w)
	err = errors.Unwrap(err)
	assert.Equal(t, orig, err)
}
