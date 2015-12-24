package main

import (
	"encoding/json"
	"fmt"
	"github.com/drone/routes"
	"github.com/xuzhenglun/container/sdk"
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

func download(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/", http.FileServer(http.Dir(Upload_Dir))).ServeHTTP(w, r)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("username") != admin {
		log.Println(r.Header.Get("username"))
		tellUserSomethingWrong(w, sdk.HaveNoRright)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	uuname := strconv.Itoa(int(time.Now().Unix())) + handler.Filename

	f, err := os.OpenFile("./upload/"+uuname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		tellUserSomethingWrong(w, sdk.ServerHandleError)
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	tellUserSuccess(w, sdk.PostSuccess, uuname)
}

func delete(w http.ResponseWriter, r *http.Request) {
	if usr := r.Header.Get("username"); usr != admin {
		tellUserSomethingWrong(w, sdk.HaveNoRright)
		return
	}
	rm := r.URL.Path
	os.Remove(Upload_Dir + rm)
	tellUserSuccess(w, sdk.DeleteSuccsess, rm)
}

func tellUserSomethingWrong(w http.ResponseWriter, c string) {
	r := sdk.Reply{Code: c}
	reply, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s", reply)
	fmt.Fprint(w, reply)
}

func tellUserSuccess(w http.ResponseWriter, code string, url string) {
	r := sdk.Reply{Code: code}
	r.Url = url
	reply, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s", reply)
	fmt.Fprintf(w, "%s", reply)
}

func main() {
	mux := routes.New()
	mux.Get("/", download)
	mux.Post("/", upload)
	mux.Del("/:items", delete)
	mux.Static("/", Upload_Dir)
	http.Handle("/", mux)
	http.ListenAndServe("localhost:8080", nil)
}
