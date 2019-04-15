package main

import (
	"flag"
	"log"
	"net/http"
	"verbio/apigw"
)

func main() {
	var (
		httpAddr   = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
		consulAddr = flag.String("consul.addr", "localhost:8500", "Consul agent address")
	)
	flag.Parse()

	api := apigw.NewAPI(*consulAddr)

	log.Printf("Listening at %s", *httpAddr)
	http.ListenAndServe(*httpAddr, api)
}
