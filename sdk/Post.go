package sdk

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func PostNewFile(fd *os.File, url string, username string) []byte {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile("file", fd.Name())
	if err != nil {
		log.Println(err)
		return nil
	}
	_, err = io.Copy(fw, fd)
	if err != nil {
		log.Println(err)
		return nil
	}
	w.Close()

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		log.Println(err)
		return nil
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client

	if username != "" {
		req.Header.Add("username", username)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}

	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return msg
	}
}
