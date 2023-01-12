package requests

import (
	"atcscraper/src/env"
	logging "atcscraper/src/log"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
	"net/url"
)

func BuildProxyClient() *retryablehttp.Client {

	ProxyEndpoint := env.LoadEnv("ZYTE_ENDPOINT")
	ProxyApiKey := env.LoadEnv("ZYTE_API_KEY")

	ProxyCert, ProxyCertLoadError := ioutil.ReadFile("static/proxy/zyte-proxy-ca.cer")
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
	RetryClient.RetryMax = 10
	RetryClient.Logger = nil

	RetryClient.HTTPClient = ProxyClient

	return RetryClient

}

func MakeGetRequestRAW(RequestURL string) string {

	ProxyClient := BuildProxyClient()

	Response, RequestError := ProxyClient.Get(RequestURL)

	if RequestError != nil {
		Error := fmt.Sprintf("Error Making RAW Request: %v", RequestError.Error())
		logging.NewError(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return string(bytes.Replace(ResponseBody, []byte("\r"), []byte("\r\n"), -1))

}

func MakeGetRequestJSON(RequestURL string) []byte {

	ProxyClient := BuildProxyClient()

	Response, RequestError := ProxyClient.Get(RequestURL)

	if RequestError != nil {
		Error := fmt.Sprintf("RequestError Making Get Request: %v", RequestError.Error())
		logging.NewError(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return ResponseBody

}