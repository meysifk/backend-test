package network

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Do(httpClient *http.Client, method, url string, queryParams map[string]string, data []byte, header map[string]string) ([]byte, error) {
	log.Printf("http request: %s; data: %s", url, string(data))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	log.Printf("http request: %s; url data: %s, header data: %+v", url, q.Encode(), header)

	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()
	code := resp.StatusCode / 100

	switch code {
	case 4:
		body, err := ioutil.ReadAll(resp.Body)
		log.Printf("http response: %s; data: %s", url, string(body))
		if err != nil {
			log.Print()
			return nil, err
		}
		msg := fmt.Sprintf("err bad request (4xx) for url [%s] -- message body [%s]", url, string(body))
		return body, errors.New(msg)
	case 5:
		body, err := ioutil.ReadAll(resp.Body)
		log.Printf("http response: %s; data: %s", url, string(body))
		if err != nil {
			log.Print(err)
			return nil, err
		}
		msg := fmt.Sprintf("err internal error (5xx) for url [%s] -- message body [%s]", url, string(body))
		return body, errors.New(msg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("http response: %s; data: %s", url, string(body))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return body, nil
}
