package phonenumbers

import (
	"testing"
)

var (
	testDataLongRange    = "02080668800-02080668850"
	testDataShortRange   = "02080668800-950"
	testDataNumberString = "02080668866"
	testDataNumberInt    = 442080668866
)

func TestSplitRange(t *testing.T) {

	firstresult, secondresult := SplitRange(testDataLongRange, "-")

	if firstresult != 442080668800 {
		t.Error("Our result should be 442080668800, got", firstresult)
	}

	if secondresult != 442080668850 {
		t.Error("Our result should be 442080668800, got", firstresult)
	}

}

func TestRetrieveTrailingDigits(t *testing.T) {

	result := RetrieveTrailingDigits(testDataNumberInt, 3)

	if result != "866" {
		t.Error("Our result should be 866, got", result)
	}

}

func TestDigitCount(t *testing.T) {

	result := DigitCount(testDataNumberInt)

	if result != 12 {
		t.Error("Our result should be 12, got", result)
	}
}

func TestConverttoE164int(t *testing.T) {

	result := ConverttoE164int(testDataNumberString)

	if result != 442080668866 {
		t.Error("Our result should be 442080668866, got", result)
	}
}

func TestConverttoE164(t *testing.T) {

	result := ConverttoE164(testDataNumberString)

	if result != "442080668866" {
		t.Error("Our result should be 442080668866, got", result)
	}

}
