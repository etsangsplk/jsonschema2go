package enum_test

import (
	"github.com/jwilner/jsonschema2go/pkg/testharness"
	"testing"
)

func TestPlan(t *testing.T) {
	testharness.RunGenerateTests(
		t,
		"testdata/",
		"testdata/generate",
		"github.com/jwilner/jsonschema2go/internal/enum/testdata",
	)
}

func TestValidation(t *testing.T) {
	testharness.RunValidationTest(t, "testdata/validation/")
}