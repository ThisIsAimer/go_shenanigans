package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type Person struct{
	Name string `json:"name"`
	Age int32 `json:"age"`
}

var personData = map[string]Person{
	"1" : {Name: "Sunpreet", Age: 20},
	"2" : {Name: "Rudra", Age: 30},
	"3" : {Name: "Shiva", Age: 40},
}

func getPersonHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		http.Error(w,"id is missing", http.StatusBadRequest)
		return
	}

	person , exists := personData[id]
	if !exists {
		http.Error(w,"id is invalid", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "Application/json")

	//posting to writer in json format
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		http.Error(w,"failed to encode response", http.StatusInternalServerError)
	}
}


func main(){
	
	http.HandleFunc("/person", getPersonHandler)

	port := ":3000"

	fmt.Println("listening to port:", port)

	cert := `certification\certificate.pem`
	key := `certification\key.pem`

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr: port,
		Handler: nil,
		TLSConfig: tlsConfig,

	}

	http2.ConfigureServer(server,&http2.Server{})

	err := server.ListenAndServeTLS(cert,key)
	
	if err != nil {
		fmt.Println("error finding port:", err)
		return
	}

}