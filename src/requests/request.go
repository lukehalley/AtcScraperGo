package requests

import (
	"atcscraper/src/env"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func BuildProxyClient() *http.Client {

	ProxyEndpoint := env.LoadEnv("ZYTE_ENDPOINT")
	ProxyApiKey := env.LoadEnv("ZYTE_API_KEY")

	ProxyCert, ProxyCertLoadError := ioutil.ReadFile("static/proxy/zyte-proxy-ca.cer")
	if ProxyCertLoadError != nil {
		log.Fatal(ProxyCertLoadError)
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

	return ProxyClient
}

func MakeGetRequestRAW(RequestURL string) string {

	ProxyClient := BuildProxyClient()

	Response, Error := ProxyClient.Get(RequestURL)

	if Error != nil {
		log.Fatalln(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return string(bytes.Replace(ResponseBody, []byte("\r"), []byte("\r\n"), -1))

}

func MakeGetRequestJSON(RequestURL string) []byte {

	ProxyClient := BuildProxyClient()

	Response, Error := ProxyClient.Get(RequestURL)

	if Error != nil {
		log.Fatalln(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return ResponseBody

}