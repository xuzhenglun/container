package main

import (
	"log"
	"net/http"
)

func main() {
	rm, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/1450880006helloworld.go", nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(rm.Method)
	client := &http.Client{}
	client.Do(rm)
}
