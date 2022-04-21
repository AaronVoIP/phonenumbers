package phonenumbers

import (
	"testing"
)

func TestSplitRange(t *testing.T) {

	exampleinput := "02080668800-02080668850"

	firstresult, secondresult := SplitRange(exampleinput, "-")

	if firstresult != 442080668800 {
		t.Error("Our result should be 442080668800, got", firstresult)
	}

	if secondresult != 442080668850 {
		t.Error("Our result should be 442080668800, got", firstresult)
	}

}
