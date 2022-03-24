package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(res.Body)
	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)
}
