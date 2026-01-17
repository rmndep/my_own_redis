package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Reader struct {
	r *bufio.Reader
}

func NewReader(conn io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(conn),
	}
}
func (parser *Reader) readLine() (line []byte, n int, err error) {

	for {
		b, err := parser.r.ReadByte()

		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)

		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}

	}

	return line[:len(line)-2], n, nil
}

func (parser *Reader) readInteger() (int, int, error) {
	line, n, err := parser.readLine()

	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)

	return int(i64), n, err
}

func (parser *Reader) readArray() (Value, error) {

	v := Value{}
	v.Typ = "array"

	size, _, err := parser.readInteger()

	if err != nil {
		return Value{}, err
	}

	v.Array = make([]Value, size)

	for i := 0; i < size; i++ {
		val, err := parser.Read()

		if err != nil {
			return Value{}, err
		}

		v.Array[i] = val
	}

	fmt.Printf("Captured Array value: %v\n", v)

	return v, nil

}

func (parser *Reader) readBulk() (Value, error) {
	v := Value{}

	v.Typ = "bulk"

	size, _, err := parser.readInteger()

	if err != nil {
		return Value{}, err
	}

	bulk := make([]byte, size)

	io.ReadFull(parser.r, bulk)

	v.Str = string(bulk)

	parser.readLine()

	return v, nil
}

// Main Read function utilize by server
func (parser *Reader) Read() (Value, error) {
	_type, err := parser.r.ReadByte()

	if err != nil {
		fmt.Println("Error occured", err)
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return parser.readArray()
	case BULK:
		return parser.readBulk()
	default:
		fmt.Printf("Unknown Type: %v\n", _type)
		return Value{}, nil
	}
}
