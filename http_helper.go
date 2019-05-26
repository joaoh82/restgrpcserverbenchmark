package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var client http.Client

func init() {
	client = http.Client{}
}

func get(path string, output interface{}) error {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Println("error creating request. ", err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("error executing request. ", err)
		return err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response body.", err)
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		log.Println("error unmarshaling response.", err)
		return err
	}

	return nil
}

func post(path string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		log.Println("error marshalling input ", err)
		return err
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		log.Println("error creating request ", err)
		return err
	}

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Println("error executing request ", err)
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response body ", err)
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		log.Println("error unmarshalling response ", err)
		return err
	}

	return nil
}

// This code was taken from https://posener.github.io/http2/
func createTLSConfigWithCustomCertificate() *tls.Config {
	// Create a pool with the server certificate since it is not signed by a CA
	caCert, err := ioutil.ReadFile("./cert/server.crt")
	if err != nil {
		log.Fatalf("error reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	return &tls.Config{
		RootCAs: caCertPool,
	}
}
