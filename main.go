package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	randomString := pb.Random{RandomString: "a_random_string", RandomInt: 42}
	bytes, err := json.Marshal(&randomString)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	server := &http.Server{Addr: "localhost:9191", Handler: http.HandlerFunc(handler)}
	log.Fatal(server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key"))
}
