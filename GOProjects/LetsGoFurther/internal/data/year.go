package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidYearFormat = errors.New("invalid year format")

type Year int32

func (r Year) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d BC", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
func (r *Year) UnmarshalJSON(jsonValue []byte) error {
	// We expect that the incoming JSON value will be a string in the format
	// "<Year> mins", and the first thing we need to do is remove the surrounding
	// double-quotes from this string. If we can't unquote it, then we return the
	// ErrInvalidYearFormat error.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidYearFormat
	}
	// Split the string to isolate the part containing the number.
	parts := strings.Split(unquotedJSONValue, " ")
	// Sanity check the parts of the string to make sure it was in the expected format.
	// If it isn't, we return the ErrInvalidYearFormat error again.
	if len(parts) != 2 || parts[1] != "BC" {
		return ErrInvalidYearFormat
	}
	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidYearFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidYearFormat
	}
	// Convert the int32 to a Year type and assign this to the receiver. Note that we
	// use the * operator to deference the receiver (which is a pointer to a Year
	// type) in order to set the underlying value of the pointer.
	*r = Year(i)
	return nil
}
