package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"heim/backend"
	"heim/backend/persist"
)

var addr = flag.String("http", ":8080", "")
var psql = flag.String("psql", "psql", "")
var static = flag.String("static", "", "")

func main() {
	flag.Parse()

	b, err := persist.NewBackend(*psql)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	server := backend.NewServer(b, *static)
	fmt.Printf("serving on %s\n", *addr)
	http.ListenAndServe(*addr, server)
}
