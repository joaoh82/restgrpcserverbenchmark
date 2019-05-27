package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
)

func getHandler(w http.ResponseWriter, _ *http.Request) {
	randomString := pb.Random{RandomString: "a_random_string", RandomInt: 42}
	bytes, err := json.Marshal(&randomString)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var random pb.Random
	if err := decoder.Decode(&random); err != nil {
		panic(err)
	}
	random.RandomString = "[Updated] " + random.RandomString

	bytes, err := json.Marshal(&random)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	log.Println("Go Backend: { HTTPVersion = 1 and 2 }; serving on https://localhost:9191/")
	log.Fatal(http.ListenAndServeTLS(":9191", "./cert/server.crt", "./cert/server.key", nil))
}
