package lib

import (
	"testing"
)

func Test_BasicSet(t *testing.T) {
	set := NewSet[string]()
	if !set.Add("Flavio") {
		t.Error("Flavio should be a new value!")
	}

	if !set.Add("Jose") {
		t.Error("Jose should be a new value!")
	}

	if set.Add("Flavio") {
		t.Error("Flavio should NOT be a new value!")
	}
}
