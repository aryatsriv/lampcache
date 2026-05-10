package parser

import "errors"

func Decode(data []byte) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}
	value, _, err := decodeOne(data)

	return value, err

}

func decodeOne(data []byte) (any, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}
	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		return readError(data)
	case ':':
		return readInt64(data)
	case '$':
		return readBulkString(data)
	case '*':
		return readArray(data)
	}

	return nil, 0, nil
}

func readArray(data []byte) (any, int, error) {
	panic("unimplemented")
}

func readBulkString(data []byte) (any, int, error) {
	panic("unimplemented")
}

func readInt64(data []byte) (any, int, error) {
	panic("unimplemented")
}

func readError(data []byte) (any, int, error) {
	panic("unimplemented")
}

func readSimpleString(data []byte) (any, int, error) {
	panic("unimplemented")
}
