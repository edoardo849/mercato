package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
)

// Price is a custom type with a precision of 2 digits
type Price float64

// Round is a custom implementation for rounding values as
// Golang does not include a round function in the standard math package
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// JSONMarshal is a custom Marshal in order to overcome the default
// behavior of the JSON encoder
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

// MarshalJSON for Price is a custom implementation
// of the Marshaler Interface
func (n Price) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}
