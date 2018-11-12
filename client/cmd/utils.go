package cmd

import (
	"io/ioutil"
	"net/http"
	"time"
)

func getBytes(url string) (bytes []byte, err error) {
	client := http.Client{Timeout: 900 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	return
}
