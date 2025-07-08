package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct{
	Name string `json:"name"`
	Age int32 `json:"age"`
}

var personData = map[string]person{
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

	err := http.ListenAndServe(port,nil)
	
	if err != nil {
		fmt.Println("error finding port:", err)
		return
	}

}


// for benchmarking wrk -t8 -c400 -d30s http://localhost:3000/person?id=2