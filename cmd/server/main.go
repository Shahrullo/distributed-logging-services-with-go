package main

import (
	"log"

	"github.com/Shahrullo/distributed-proglog-servers-with-go/interval/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
