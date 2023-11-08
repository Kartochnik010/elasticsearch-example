package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var (
	posturl  = "http://localhost:9200/products"
	postbody = `{
		"name": "Awesome T-Shirt", 
		"description": "This is an awesome t-shirt for casual wear.", 
		"price": 19.99, 
		"category": "Clothing", 
		"brand": "Example Brand"
	}`

	client = &http.Client{}
)

func index() http.Handler {
	b := bytes.NewBuffer([]byte(postbody))
	r, err := http.NewRequest(http.MethodPost, posturl, b)
	if err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", err)
		})
	}
	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", err)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", content)
	})

}

func main() {

	http.Handle("/", index())
	http.ListenAndServe(":8080", nil)
}
