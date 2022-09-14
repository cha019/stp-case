package sdk

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Request     *http.Request
	Verbose     bool
	RequestName string
}

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetHttpClient() *http.Client {

	var client *http.Client
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		//MaxIdleConns: 0,
		//MaxConnsPerHost: 0,
		MaxIdleConnsPerHost: 2000,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		DisableKeepAlives:   false,
	}

	client = &http.Client{
		Transport: httpTransport,
		Timeout:   time.Duration(15) * time.Second,
	}

	return client
}
