package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PingHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	expected := "pong"
	if rr.Body.String() != expected {
		t.Errorf("Unexpected body: got %q, want %q", rr.Body.String(), expected)
	}
}

func TestVersionHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	expected := fmt.Sprintf("%s v%s", Name, Version)
	if rr.Body.String() != expected {
		t.Errorf("Unexpected body: got %q, want %q", rr.Body.String(), expected)
	}
}

func TestKillHandler(t *testing.T) {
	// Here we only test if the end point kills itself.
	if os.Getenv("KILL_ENDPOINT") == "1" {
		req, err := http.NewRequest("DELETE", "/kill", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(KillHandler)
		handler.ServeHTTP(rr, req)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestKillHandler")
	cmd.Env = append(os.Environ(), "KILL_ENDPOINT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("Process ran with err %v, want exit status 1", err)
}

func TestKillHandler_WrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/kill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(KillHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusMethodNotAllowed)
	}
}
