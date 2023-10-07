package data
import (
"fmt"
"strconv"
)
type Year int32
func (r Year) MarshalJSON() ([]byte, error) {
jsonValue := fmt.Sprintf("%d BC", r)
quotedJSONValue := strconv.Quote(jsonValue)
return []byte(quotedJSONValue), nil
}