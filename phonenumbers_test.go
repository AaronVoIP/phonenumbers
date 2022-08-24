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

func TestListRangeShort(t *testing.T) {

	result := ListRange(44208066850, 7900)

	if result[5] != "44208066855" {
		t.Error("Our result should be a slice of numbers, we got", result[5])
	}

	if result[80] != "44208066930" {
		t.Error("Our result should be a slice of numbers, we got", result[80])
	}

	if result[100] != "44208066950" {
		t.Error("Our result should be a slice of numbers, we got", result[100])
	}

	if result[1000] != "44208067850" {
		t.Error("Our result should be a slice of numbers, we got", result[1000])
	}

	result = ListRange(44208066850, 877)

	if result[17] != "44208066867" {
		t.Error("Our result should be a slice of numbers, we got", result[17])
	}

	if result[25] != "44208066875" {
		t.Error("Our result should be a slice of numbers, we got", result[25])
	}

	if result[27] != "44208066877" {
		t.Error("Our result should be a slice of numbers, we got", result[27])
	}

	result = ListRange(44208066850, 77)

	if result[17] != "44208066867" {
		t.Error("Our result should be a slice of numbers, we got", result[17])
	}

	if result[25] != "44208066875" {
		t.Error("Our result should be a slice of numbers, we got", result[25])
	}

	if result[27] != "44208066877" {
		t.Error("Our result should be a slice of numbers, we got", result[27])
	}

}

func TestListRangeLong(t *testing.T) {

	result := ListRange(44208066850, 44208066877)

	if result[17] != "44208066867" {
		t.Error("Our result should be a slice of numbers, we got", result[17])
	}

	if result[25] != "44208066875" {
		t.Error("Our result should be a slice of numbers, we got", result[25])
	}

	if result[27] != "44208066877" {
		t.Error("Our result should be a slice of numbers, we got", result[27])
	}

}

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

func TestISOGeoData(t *testing.T) {

	response, _ := GeoData(44)

	for _, v := range response.([]string) {

		if v != "United Kingdom" {
			t.Errorf("Should have United Kingdom as response, got %v", response)
		}

	}

}

func TestNameGeoData(t *testing.T) {

	response, _ := GeoData("United Kingdom")

	for _, v := range response.([]int) {

		if v != 44 {
			t.Errorf("Should have 44 as response, got %v", response)
		}

	}

}

func TestNilGeoData(t *testing.T) {

	response, _ := GeoData(nil)

	if len(response.(map[string]int)) == 0 {
		t.Error("Map response is empty")
	}

}
