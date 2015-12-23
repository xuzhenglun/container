package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	admin      = "Reficul"
	Upload_Dir = "./upload"
)

type Reply struct {
	Code string
	Url  string
}

func service(w http.ResponseWriter, r *http.Request) {
	log.Println()
	switch r.Method {
	case "GET":
		download(w, r)
	case "POST":
		upload(w, r)
	case "DELET":
		fmt.Fprint(w, "have not supported")
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/", http.FileServer(http.Dir(Upload_Dir))).ServeHTTP(w, r)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("username") != admin {
		log.Println(r.Header.Get("username"))
		tellUserSomethingWrong(w, "403")
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)

	uuname := strconv.Itoa(int(time.Now().Unix())) + handler.Filename

	f, err := os.OpenFile("./upload/"+uuname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		tellUserSomethingWrong(w, "400")
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	tellUserSuccess(w, uuname)
}

func tellUserSomethingWrong(w http.ResponseWriter, c string) {
	r := Reply{Code: c}
	reply, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s", reply)
	fmt.Fprint(w, reply)
}

func tellUserSuccess(w http.ResponseWriter, url string) {
	r := Reply{Code: "200"}
	r.Url = url
	reply, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s", reply)
	fmt.Fprint(w, reply)
}

func main() {
	http.HandleFunc("/", service)
	http.ListenAndServe("localhost:8080", nil)
}
