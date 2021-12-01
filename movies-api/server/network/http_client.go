package network

import (
	"net/http"
	"time"
)

func New(httpTimeout int) *http.Client {
	return &http.Client{
		Timeout: time.Millisecond * time.Duration(httpTimeout),
	}
}
