package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	r := mySum(17000, 42000)
	if r != 59000 {
		t.Error("Expected", 59000, "Got", r)
	}
}
