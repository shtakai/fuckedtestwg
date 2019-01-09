package underscore

import "testing"

func TestCamel(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithA", "ends_with_a"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := Camel(tt.arg); got != tt.want {
				t.Fatalf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
