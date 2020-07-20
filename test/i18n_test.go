package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// TestI18n rule.
func TestI18n(t *testing.T) {
	testRule(t, "i18n", &rule.I18nRule{})
}
