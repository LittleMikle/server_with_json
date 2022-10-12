package web

import (
	"fmt"
	"net/http"
)

func RunWeb(address string) {

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(address, nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome,Your IP and port is:%s\n", r.RemoteAddr)

}
