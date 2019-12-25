package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//net /http
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		fmt.Sprintf("%v\n", err)
	}
	w.Write(b)

}
func main() {
	http.HandleFunc("/hello/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
