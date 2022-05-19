package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Dices []int
	Sum   int
}

func handler(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()
	var dices, sum, error = throw(query.Get("exp"))
	if error != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad request")
	} else {
		var response = Response{dices, sum}
		r.Header.Add("Content-Type", "application/json")
		var encoded, _ = json.Marshal(response)
		w.Write(encoded)
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
