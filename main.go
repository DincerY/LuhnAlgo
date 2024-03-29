package main

import (
	"LuhnAlgo/luhn"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/post", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Hata: %s", err)
		return
	}
	fmt.Fprintf(w, luhn.RunAlgo(string(body)))
}
