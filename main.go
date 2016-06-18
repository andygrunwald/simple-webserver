package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

const (
	// Name of the application
	Name = "simple webserver"
	// Version of the application
	Version = "1.0.0"
)

func main() {
	var (
		listenFlag = flag.String("listen", ":8082", "Address + Port to listen on. Format ip:port.")
		redisFlag  = flag.String("redis", ":6379", "Address + Port where a redis server is listening.")
	)
	flag.Parse()

	// Create Redis storage
	r := NewRedisStorage(*redisFlag)

	// Define HTTP endpoints
	s := http.NewServeMux()
	s.HandleFunc("/", RootHandler)
	s.HandleFunc("/ping", PingHandler(r))
	s.HandleFunc("/kill", KillHandler)
	s.HandleFunc("/version", VersionHandler)

	// Bootstrap logger
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Printf("Starting webserver and listen on %s", *listenFlag)

	// Start HTTP Server with request logging
	loggingHandler := handlers.LoggingHandler(os.Stdout, s)
	log.Fatal(http.ListenAndServe(*listenFlag, loggingHandler))
}

// RootHandler handles requests to the "/" path.
// It will redirect the request to /ping with a 303 HTTP header
func RootHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/ping", http.StatusSeeOther)
}

// PingHandler handles request to the "/ping" endpoint.
// It will send a PING request to Redis and return the response
// of the NoSQL database.
// The response is obvious: "pong" :)
func PingHandler(s Storage) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		res, err := s.Ping()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
			return
		}
		resp.WriteHeader(http.StatusOK)
		fmt.Fprintf(resp, res)
	}
}

// KillHandler handles request to the "/kill" endpoint.
// Will shut down the webserver immediately (via exit code 1).
// Only DELETE requests are accepted.
// Other request methods will throw a HTTP Status Code 405 (Method Not Allowed)
func KillHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We need to send a HTTP Status Code 200 (OK)
	// to respond that we have accepted the request.
	// Here we send a chunked response to the requester.
	flusher, ok := resp.(http.Flusher)
	if !ok {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusOK)
	flusher.Flush()

	// And we kill the server (like requested)
	os.Exit(1)
}

// VersionHandler handles request to the "/version" endpoint.
// It prints the Name and Version of this app.
func VersionHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "%s v%s", Name, Version)
}
