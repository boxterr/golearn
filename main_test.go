package learningGo

import (
	"testing"
)

func TestMySum(t *testing.T) {
	r := mySum(5, 10)
	if r != 15 {
		t.Error("Expected", 15, "Got", r)
	}
}
