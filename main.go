package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	// i want to use my local fmt package tho
	_, err := fmt.Fprintf(writer, "blorg")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var (
		err  error
		addr string
	)
	addr = ":9999"
	http.HandleFunc("/", handler)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		// this'll never be nil tho
		panic(err.Error())
	}
}
