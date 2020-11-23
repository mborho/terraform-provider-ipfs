package main

import (
	"crypto/tls"
	//"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type PinService struct {
	name  string
	host  string
	token string
}

type RemotePinClient struct {
	client   *http.Client
	services interface{}
}

func buildPinServiceConfigs(configs []interface{}) map[string]*PinService {
	if len(configs) < 1 {
		return nil
	}

	services := make(map[string]*PinService)
	for _, config := range configs {
		c := config.(map[string]interface{})
		service := &PinService{
			name:  c["name"].(string),
			host:  c["host"].(string),
			token: c["token"].(string),
		}
		services[service.name] = service
	}
	return services
}

/*func NewRemotePinClient(services []string) (*RemotePinClient, error) {
	log.Println("[INFO] Building remote pin client")

	// build transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// detect https proxy in environment
	proxy_https := os.Getenv("HTTPS_PROXY")
	if proxy_https == "" {
		proxy_https = os.Getenv("https_proxy")
	}

	// check proxy
	if proxy_https != "" {
		proxyUrl, err := url.Parse(proxy_https)
		if err != nil {
			return nil, err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	// build http client
	var netClient = &http.Client{
		Timeout:   time.Minute * 10,
		Transport: tr,
	}

	// return client
	client := &RemotePinClient{
		client:   netClient,
		services: services,
	}
	return client, nil
}*/
