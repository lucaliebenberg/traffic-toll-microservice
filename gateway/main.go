package main

import "net/http"

func main() {
	http.HandleFunc("/invoice", handleGetInvoice)
	http.ListenAndServe(":6000", nil)
}

func handleGetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is working just fine"))
}
