package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func userDetails(w http.ResponseWriter, req *http.Request) {
	jsonData, err := ioutil.ReadFile("storage/user_data.json")
	if err != nil {
		http.Error(w, "Failed to read JSON file", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write JSON response:", err)
	}
}
func userDetailsP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	jsonData, err := ioutil.ReadFile("storage/user_data.json")
	if err != nil {
		http.Error(w, "Failed to read JSON file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write JSON response:", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/getuserdetails", userDetails)
	http.HandleFunc("/getuserdetailsp", userDetailsP)

	http.ListenAndServe(":8090", nil)
}
