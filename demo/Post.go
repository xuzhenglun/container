package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	admin = "Reficul"
	URL   = "http://127.0.0.1:8080"
)

func main() {

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile("file", "helloworld.go")
	if err != nil {
		log.Println(err)
		return

	}
	fd, err := os.Open("examplefile")
	if err != nil {
		log.Println(err)
		return

	}
	defer fd.Close()

	_, err = io.Copy(fw, fd)
	if err != nil {
		log.Println(err)
		return

	}
	w.Close()
	req, err := http.NewRequest("POST", URL, buf)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client

	req.Header.Add("username", admin)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	io.Copy(os.Stdout, res.Body)
}
