package underscore

import "testing"

func TestCamel_simple(t *testing.T) {
	arg := "thisIsACamelCaseString"
	want := "this_is_a_camel_case_string"
	got := Camel(arg)

	if got != want {
		t.Errorf("Camel(%q) = %q; want %q", arg, got, want)
	}
}
