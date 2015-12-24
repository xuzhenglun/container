package sdk

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Deletefile(file string, username string) []byte {
	rm, err := http.NewRequest("DELETE", file, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	rm.Header.Add("username", username)
	client := &http.Client{}
	res, err := client.Do(rm)
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return msg
}
