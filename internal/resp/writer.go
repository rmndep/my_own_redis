package resp

import (
	"bufio"
	"fmt"
	"io"
)

type Writer struct {
	w *bufio.Writer
}

func NewWriter(conn io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(conn),
	}
}

func (writer *Writer) write(v Value) []byte {

	switch v.Typ {

	case "error":
		return []byte(fmt.Sprintf("-%s\r\n", v.Str))

	case "string":
		return []byte(fmt.Sprintf("+%s\r\n", v.Str))

	case "bulk":
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v.Str), v.Str))

	case "array":

		res := []byte(fmt.Sprintf("*%d\r\n", len(v.Array)))

		for _, item := range v.Array {
			bytes := writer.write(item)
			res = append(res, bytes...)
		}

		return res

	default:
		fmt.Println("Unsupported type: %v", v)
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", 5, "error"))
	}
}

func (writer *Writer) Write(v Value) error {

	fmt.Printf("Writer %v\n", v)

	bytes := writer.write(v)

	_, err := writer.w.Write(bytes)

	fmt.Println("sending", string(bytes))

	if err != nil {
		fmt.Println("Error writing back", err)
		return err
	}

	return writer.w.Flush()
}
