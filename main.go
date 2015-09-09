package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8082", "Address + Port to listen on. Format ip:port.")
	)
	flag.Parse()

	log.Printf("Starting webserver and listen on %s", *listen)
	http.HandleFunc("/", httpHandleRoot)
	http.HandleFunc("/ping", httpHandlePing)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func httpHandleRoot(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/ping", http.StatusSeeOther)
}

func httpHandlePing(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintln(resp, "pong")
}
