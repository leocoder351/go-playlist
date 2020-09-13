package string_fn

import (
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for a, part := range parts {
		t.Log(a)
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}
