package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	Name = "simple webserver"
	Version = "0.1.0-dev"
)

func main() {
	var (
		listen = flag.String("listen", ":8082", "Address + Port to listen on. Format ip:port.")
	)
	flag.Parse()

	http.HandleFunc("/", httpHandleRoot)
	http.HandleFunc("/ping", httpHandlePing)
	http.HandleFunc("/version", VersionHandler)

	log.Printf("Starting webserver and listen on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func httpHandleRoot(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/ping", http.StatusSeeOther)
}

func httpHandlePing(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintln(resp, "pong")
}

func VersionHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "%s v%s", Name, Version)
}