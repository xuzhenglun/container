package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuzhenglun/container/sdk"
	"log"
	"os"
)

const (
	admin = "Reficul"
	URL   = "http://127.0.0.1:8080"
)

func main() {

	filename := "examplefile"
	if len(os.Args) > 1 && os.Args[1] != "" {
		filename = os.Args[1]
	}
	fd, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	var reply sdk.Reply
	if msg := sdk.PostNewFile(fd, URL, admin); msg != nil {
		json.Unmarshal(msg, &reply)
		fmt.Printf("Response: %s\n", string(msg))
		fmt.Printf("Return code: %s, file name in server is %s\n", reply.Code, reply.Url)
	} else {
		fmt.Println("Unkown Err")
		return
	}

	if reply.Code == "200" {
		fmt.Println("Press any key to Delete file")

		var s string
		fmt.Scanln(&s)

		msg := sdk.Deletefile(URL+"/"+reply.Url, admin)
		if msg != nil {
			json.Unmarshal(msg, &reply)
			fmt.Printf("Response: %s\n", string(msg))
			fmt.Printf("Return code: %s, file name in server is %s\n", reply.Code, reply.Url)
		} else {
		}
	}
}
