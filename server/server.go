package main

import (
	"net/http"
	"net/rpc"
)

func main() {

	mit := NewVegetableController()

	rpc.Register(mit)

	rpc.HandleHTTP()

	http.ListenAndServe(":9000", nil)

}
