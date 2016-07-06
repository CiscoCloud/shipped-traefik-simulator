// Simple API server
// Test Server api to generate load and return cutom http status code and
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", page)
	http.HandleFunc("/hits", hitServer)
	port := ":9090"
	log.Printf("Server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

//Request input request parameters
type Request struct {
	URL        string
	StatusCode int
	Delay      int
}

func page(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func hitServer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error, Invalid Request")
	}
	reg := new(Request)
	if err = json.Unmarshal(b, reg); err != nil {
		fmt.Fprintf(w, "Error, Failled to decode Request err: %s", err)
	}

	time.Sleep(time.Duration(reg.Delay) * time.Millisecond)

	if reg.StatusCode != 0 {
		w.WriteHeader(reg.StatusCode)
	}
	fmt.Fprint(w, "Done ")

}
