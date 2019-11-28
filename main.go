package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := ""
	if r.Body != nil {
		str, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		body = string(str)
	}

	log.Printf("%s %s %s", r.Method, r.URL.Path, body)

	fmt.Fprint(w, "Success")

}
