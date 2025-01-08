package requests

import (
	"atcscraper/src/env"
	logging "atcscraper/src/log"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"os"
	"net/url"
)

func BuildProxyClient(RetryCount int) *retryablehttp.Client {

	ProxyEndpoint := env.LoadEnv("ZYTE_ENDPOINT")
	ProxyApiKey := env.LoadEnv("ZYTE_API_KEY")

	ProxyCert, ProxyCertLoadError := os.ReadFile("static/proxy/zyte-proxy-ca.cer")
	if ProxyCertLoadError != nil {
		Error := fmt.Sprintf("Error Loading ZYTE Proxy Cert: %v", ProxyCertLoadError.Error())
		logging.NewError(Error)
	}

	ProxyCertPool := x509.NewCertPool()
	ProxyCertPool.AppendCertsFromPEM(ProxyCert)

	RawURL := "http://" + ProxyApiKey + ":@" + ProxyEndpoint
	ProxyURL, _ := url.Parse(RawURL)
	ProxyClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(ProxyURL),
			TLSClientConfig: &tls.Config{
				RootCAs: ProxyCertPool,
			},
		},
	}

	RetryClient := retryablehttp.NewClient()
	RetryClient.StandardClient()
	RetryClient.RetryMax = RetryCount
	RetryClient.Logger = nil

	RetryClient.HTTPClient = ProxyClient

	return RetryClient

}

func MakeGetRequestRAW(RequestURL string, RetryCount int) string {

	ProxyClient := BuildProxyClient(RetryCount)

	Response, RequestError := ProxyClient.Get(RequestURL)

	if RequestError != nil {

	}

	if RequestError != nil {
		// Error := fmt.Sprintf("Warning: Error Making RAW Request: %v", RequestError.Error())
		// log.Print(Error)
		return ""
	} else {
		ResponseBody, _ := io.ReadAll(Response.Body)
		return string(bytes.Replace(ResponseBody, []byte("\r"), []byte("\r\n"), -1))
	}



}

func MakeGetRequestJSON(RequestURL string, RetryCount int) []byte {

	ProxyClient := BuildProxyClient(RetryCount)

	Response, RequestError := ProxyClient.Get(RequestURL)

	var EmptyByteArray []byte
	if RequestError != nil {
		// Error := fmt.Sprintf("RequestError Making Get Request: %v", RequestError.Error())
		// log.Print(Error)
		return EmptyByteArray
	} else {
		ResponseBody, _ := io.ReadAll(Response.Body)
		return ResponseBody
	}

}