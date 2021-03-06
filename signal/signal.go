package signal

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Age        int
	Name       string
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Age:        21,
		Name:       "Bob Sapp",
		Occupation: "Nurse",
	}
	data, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Internal fucked up error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
