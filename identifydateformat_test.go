package identifydateformat

import (
	"reflect"
	"testing"
)

// YEAR: 2006
// MONTH: 01
// DAY: 02

// TestIdentifyDateFormat returns the only format that matches the given dates
func TestIdentifyDateFormat(t *testing.T) {
	expect := []string{"02/01/06"}
	got, err := IdentifyDateFormat([]string{"02/02/02", "23/12/01", "10/12/32", "01/10/01"})
	if err != nil {
		t.Errorf("Expected '%#v', got nil\n", expect)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected '%#v', got '%#v'\n", expect, got)
	}
}

// TestErrorIdentifyingDateFormat should return error as there is no single
// format that could matches the given dates
func TestErrorIdentifyingDateFormat(t *testing.T) {
	got, err := IdentifyDateFormat([]string{"30/03/1980", "03/30/1980", "1980-03-30"})
	if err == nil {
		t.Errorf("Expected error, got '%#v'\n", got)
	}
}

// TestAddPossibleFormats do not add any formats that already exists
func TestAddPossibleFormats(t *testing.T) {
	DateFormats = []string{"a", "b", "c"}
	AddPossibleFormats("b", "a", "d", "a", "d")

	expect := []string{"a", "b", "c", "d"}
	got := DateFormats
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected '%#v', got '%#v'\n", expect, got)
	}
}
