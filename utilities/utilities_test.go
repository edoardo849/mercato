package utilities

import (
	"bytes"
	"testing"
)

func TestRound(t *testing.T) {

	var floatTests = []struct {
		val     float64 // input
		roundOn float64
		places  int
		e       float64 // expected result
	}{
		{34.5678, 0.5, 2, 34.57},
		{12.3456787, 0.5, 3, 12.346},
		{12.123123, 0.5, 4, 12.1231},
	}

	for _, tt := range floatTests {
		actual := Round(tt.val, tt.roundOn, tt.places)
		if actual != tt.e {
			t.Errorf("Rounding(%f): expected %f, actual %f", tt.val, tt.e, actual)
		}
	}
}

func TestJSONMarshal(t *testing.T) {

	type testStruct struct {
		X string
	}

	testVal := testStruct{
		X: "& test & test < >",
	}

	res2B, err := JSONMarshal(testVal, true)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	res2C, err := JSONMarshal(testVal, false)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if bytes.Contains(res2B, []byte("\\u003c")) {
		t.Errorf("Should not contain \\u003c but <")
	}

	if bytes.Contains(res2B, []byte("\\u003e")) {
		t.Errorf("Should not contain \\u003e but >")
	}

	if bytes.Contains(res2B, []byte("\\u0026")) {
		t.Errorf("Should not contain \\u003c but &")
	}

	if !bytes.Contains(res2C, []byte("\\u003c")) {
		t.Errorf("Should contain \\u0026, not <")
	}

	if !bytes.Contains(res2C, []byte("\\u003e")) {
		t.Errorf("Should contain \\u003e, not >")
	}

	if !bytes.Contains(res2C, []byte("\\u0026")) {
		t.Errorf("Should contain \\u0026, not &")
	}

}
