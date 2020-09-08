package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	httpHandler()
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpHandler() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Running..."))
		if err != nil {
			log.Println("Error", err)
		}
	})
	http.HandleFunc("/lebronjames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(callNBAAPI("Lebron"))
		if err != nil {
			log.Println("Error", err)
		}
	})
	http.HandleFunc("/jamesharden", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(callNBAAPI("Harden"))
		if err != nil {
			log.Println("Error", err)
		}
	})
	http.HandleFunc("/getPlayers", func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Player"]) == 0 {
			log.Println("Error - missing player Header")
			w.WriteHeader(404)
		} else {
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(callNBAAPI(r.Header["Player"][0]))
		if err != nil {
			log.Println("Error", err)
		}
	}
	})
}

func callNBAAPI(player string) []byte {
	response, err := http.Get(fmt.Sprintf("https://www.balldontlie.io/api/v1/players?search=%s", player))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}