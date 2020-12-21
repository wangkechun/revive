package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// TestErrWrap rule.
func TestErrWrap(t *testing.T) {
	testRule(t, "err-wrap", &rule.ErrWrapRule{})
}
