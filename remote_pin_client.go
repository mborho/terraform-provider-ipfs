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
	name          string
	host          string
	token         string
	skipSSLVerify bool
	client        *http.Client
}

type RemotePinClient struct {
	services interface{}
}

func buildPinServiceConfigs(configs []interface{}) map[string]*PinService {
	if len(configs) < 1 {
		return nil
	}

	services := make(map[string]*PinService)
	log.Println("[INFO] Load remote service remote clients.")
	for _, config := range configs {
		c := config.(map[string]interface{})
		service := &PinService{
			name:          c["name"].(string),
			host:          c["host"].(string),
			token:         c["token"].(string),
			skipSSLVerify: c["skip_ssl_verify"].(bool) == false,
		}
		if err := service.loadHttpClient(); err != nil {
			log.Println("[INFO] Load remote service remote clients failed: ", err)
		}
		services[service.name] = service
	}
	return services
}

func (p *PinService) loadHttpClient() error {

	// build transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: p.skipSSLVerify},
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
			return err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	// build http client
	p.client = &http.Client{
		Timeout:   time.Minute * 10,
		Transport: tr,
	}
	return nil
}

func (p *PinService) AddPin(cid, name string, origins []interface{}) error {
	//req.Header("Authorization", fmt.Sprintf("Bearer %s", client.temporalToken))
	return nil
}
