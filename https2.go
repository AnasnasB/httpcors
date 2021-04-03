package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Lenght, Authorization")
		return
	}
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return 
		}
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	}
	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are on page1")
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/page1", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
