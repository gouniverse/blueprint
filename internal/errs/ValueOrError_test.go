package errs

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_ValueOrErrorSuccess(t *testing.T) {
	initial := "zero"

	str, err := ValueOrError(func() (string, error) {
		initial += " one"
		return initial, nil
	}, func() (string, error) {
		initial += " two"
		return initial, nil
	})

	if err != nil {
		t.Fatal(`unexpected error`, err)
	}

	if str != "zero one two" {
		t.Fatal(`str should be "zero one two", but found:`, str)
	}
}

func Test_ValueOrErrorFail(t *testing.T) {
	initial := "zero"

	str, err := ValueOrError(func() (string, error) {
		initial += " one"
		return initial, nil
	}, func() (string, error) {
		return "failed on purpose", errors.New("failed on purpose")
	})

	if err == nil {
		t.Fatal(`must return error`)
	}

	if str != "failed on purpose" {
		t.Fatal(`str should be "failed on purpose", but found:`, str)
	}
}
