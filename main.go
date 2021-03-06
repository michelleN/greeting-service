package main

import (
	"encoding/json"
	"net/http"
)

// Greeting represents what is returned in the response
type Greeting struct {
	Message string
	Color   string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{"Good Night", "#00008b"}

	js, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
