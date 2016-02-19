package faithplusone

import "testing"

func TestYourFaith(t *testing.T) {
	faith := 7
	plusOne := FaithPlusOne(faith)
	if plusOne != 8 {
		t.Errorf("Expected 8 got %v", plusOne)
	}
}
