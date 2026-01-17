package main

import (
	"fmt"
	"net"
	"os"
	"redis/internal/resp"
	"runtime"
)

func handler(v resp.Value) resp.Value {

	command := v.Array[0].Str

	handlerFunc := Dispatcher(command)

	val := handlerFunc(v.Array[1:])

	return val
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	reader := resp.NewReader(conn)
	writer := resp.NewWriter(conn)

	fmt.Println("New Client! Active Goroutines:", runtime.NumGoroutine())

	for {

		v, err := reader.Read()

		if err != nil {
			fmt.Println("Error reading conn data", err)
			return
		}

		response := handler(v)

		fmt.Println("RESP", response)

		writer.Write(response)
	}

}

func main() {

	fmt.Println("Listening on PORT 6379")

	ls, err := net.Listen("tcp", ":6379")

	defer ls.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := ls.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)

	}
}
