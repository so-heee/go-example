package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Ping struct {
	Status int
	Result string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok"}
	res, _ := json.Marshal(ping)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(res)
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", rootHandler)
	httpServer.Addr = ":8888"
	log.Println(httpServer.ListenAndServe())
}
