package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomResponse struct {
	Document string `json:"document"`
}

const port = 8001

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", HelloServer2)
	fmt.Println("running on ", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	fmt.Fprint(w, "Hello-test!")
}

func HelloServer2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	w.Header().Set("Content-Type", "application/json")
	res := CustomResponse{Document: "new doc"}
	json.NewEncoder(w).Encode(res)
}
