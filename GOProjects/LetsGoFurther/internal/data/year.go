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

	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidYearFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "BC" {
		return ErrInvalidYearFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidYearFormat
	}

	*r = Year(i)
	return nil
}
