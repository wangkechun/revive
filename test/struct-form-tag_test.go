package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// TestStructFormTag rule.
func TestStructFormTag(t *testing.T) {
	testRule(t, "struct-form-tag", &rule.StructFormTagRule{})
}
