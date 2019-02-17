package api

import (
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
)

var PfxApi = "http://5c66d40524e2140014f9edfb.mockapi.io/api"

func NewApiRequest(uri, method string, body interface{}) (code int, data []byte) {
	var req *http.Request
	var err error

	client := &http.Client{}
	client.Timeout = time.Second * 15

	uri = PfxApi + uri

	if body != nil {
		d, err := json.Marshal(body)
		if err != nil {
			fmt.Println("###ERR: NewApiRequest->json.Marshal(body):", err)
			return
		}

		req, err = http.NewRequest(method, uri, bytes.NewBuffer(d))
	} else {
		req, err = http.NewRequest(method, uri, nil)
	}

	if err != nil {
		fmt.Println("###ERR: NewApiRequest->http.NewRequest():", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("###ERR: NewApiRequest->client.Do(req):", err)
		return
	}

	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("###ERR: NewApiRequest->ioutil.ReadAll(resp.Body):", err)
		return
	}

	return resp.StatusCode, data
}
