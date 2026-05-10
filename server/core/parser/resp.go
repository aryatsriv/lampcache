package parser

import "errors"

func Decode(data []byte) (interface{}, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("no data")
	}
	value, _, err := decodeOne(data)

	return value, err

}

func decodeOne(data []byte) (interface{}, int, error) {
	panic("unimplemented")
}
