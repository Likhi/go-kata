package main

import "testing"

func TestToCamelCase(t *testing.T) {
	s := "hello-there"
	sp := "helloThere"
	if ToCamelCase(s) != sp {
		t.Errorf("Expected %s. Got %s.", s, sp)
	}

}
