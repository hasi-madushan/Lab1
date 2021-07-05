package main

import (
	"net/http"
	"net/rpc"
)

func main() {

	mit := NewVegetableFactory()

	// register `mit` object with `rpc.DefaultServer`
	rpc.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
