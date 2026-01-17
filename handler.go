package main

import (
	"fmt"
	"redis/internal/resp"
	"strings"
)

type HandlerFunc func(args []resp.Value) resp.Value

const (
	SET  = "SET"
	GET  = "GET"
	PING = "PING"
	ECHO = "ECHO"
	OK   = "OK"
)

func Ping(args []resp.Value) resp.Value {

	if len(args) == 0 {
		return resp.Value{Typ: "string", Str: "PONG"}
	}

	if len(args) > 1 {
		return resp.Value{Typ: "error", Str: "Err wrong number of arguments for 'ping' command"}
	}

	response := outputStr(args)

	return resp.Value{Typ: "string", Str: response}

}

func Error(args []resp.Value) resp.Value {

	err := fmt.Sprintf("Unknown Command")

	return resp.Value{Typ: "error", Str: err}
}

func outputStr(args []resp.Value) string {
	size := len(args)

	strArr := make([]string, size)

	for i := 0; i < size; i++ {
		strArr[i] = args[i].Str
	}

	return strings.Join(strArr, " ")
}

func Echo(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return invalid(ECHO)
	}

	echo := outputStr(args)

	return resp.Value{Typ: "string", Str: echo}

}

func invalid(command string) resp.Value {
	return resp.Value{Typ: "error", Str: fmt.Sprintf("Err wrong number of arguments for '%s' command", command)}
}

func ok() resp.Value {
	return resp.Value{Typ: "string", Str: OK}
}

var setMap = map[string]any{
	"version": "1.0.0",
}

func Set(args []resp.Value) resp.Value {
	if len(args) == 0 || len(args) == 1 {
		return invalid(SET)
	}

	key := args[0].Str
	value := outputStr(args[1:])

	setMap[key] = value

	fmt.Println(setMap)

	return ok()
}

func Get(args []resp.Value) resp.Value {
	if len(args) > 1 {
		return invalid(GET)
	}

	key := args[0].Str

	val, ok := setMap[key].(string)

	if ok == false {
		return resp.Value{}
	}

	return resp.Value{Typ: "string", Str: val}
}

var dispatcher = map[string]HandlerFunc{
	PING: Ping,
	ECHO: Echo,
	SET:  Set,
	GET:  Get,
}

func Dispatcher(command string) HandlerFunc {

	command = strings.ToUpper(command)

	dispatch, ok := dispatcher[command]

	if ok == false {
		return Error
	}

	return dispatch
}
