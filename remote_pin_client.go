package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	pinningSuccessCodes = map[int]bool{200: true, 202: true}
)

type PinService struct {
	name          string
	endpoint      string
	token         string
	skipSSLVerify bool
	client        *http.Client
}

type RemotePinObject struct {
	Cid     string            `json:"cid"`
	Name    string            `json:"name"`
	Origins []string          `json:"origins"`
	Meta    map[string]string `json:"meta"`
}

type RemotePinResponse struct {
	RequestId string            `json:"requestid"`
	Status    string            `json:"status"`
	Created   string            `json:"created"`
	Pin       RemotePinObject   `json:"pin"`
	Delegates []string          `json:"delegates"`
	Info      map[string]string `json:"info"`
}

func buildPinServiceConfigs(configs []interface{}) map[string]*PinService {
	if len(configs) < 1 {
		return nil
	}

	services := make(map[string]*PinService)
	for _, config := range configs {
		c := config.(map[string]interface{})
		log.Printf("[INFO] Load remote service remote client %s.", c["name"])
		service := &PinService{
			name:          c["name"].(string),
			endpoint:      c["endpoint"].(string),
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

func (p *PinService) doApiRequest(mode string, apiUrl string, data []byte, successCodes map[int]bool, contentType ...string) ([]byte, error) {
	var postData io.Reader

	// handle content type
	contentTypeHeader := "application/json"
	if len(contentType) > 0 {
		contentTypeHeader = contentType[0]
	}

	log.Printf("API Payload %s\n", data)

	// handle post data
	if len(data) > 0 {
		postData = bytes.NewBuffer(data)
	}
	request, err := http.NewRequest(mode, apiUrl, postData)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentTypeHeader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.token))

	response, err := p.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// parse response
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("API RESPONSE %s [%s]\n%s\n", apiUrl, response.StatusCode, contents)

	if _, ok := successCodes[response.StatusCode]; ok == false {
		log.Printf("[INFO] STATUS CODE %s %s\n", apiUrl, response.StatusCode)
		log.Println(contents)
		return nil, errors.New(fmt.Sprintf("API request failed:%s", contents))
	}
	return contents, err
}

func (p *PinService) PostPin(apiUrl, cid, name string, origins []interface{}, meta map[string]interface{}) (*RemotePinResponse, error) {
	var originsSlice []string
	for _, param := range origins {
		originsSlice = append(originsSlice, param.(string))
	}

	metaMap := make(map[string]string)
	for k, v := range meta {
		metaMap[k] = v.(string)
	}

	pin := &RemotePinObject{
		Cid:     cid,
		Name:    name,
		Origins: originsSlice,
		Meta:    metaMap,
	}

	reqBody, err := json.Marshal(pin)
	if err != nil {
		return nil, err
	}

	contents, err := p.doApiRequest("POST", apiUrl, reqBody, pinningSuccessCodes)
	if err != nil {
		return nil, err
	}

	responseData := RemotePinResponse{}
	err = json.Unmarshal(contents, &responseData)
	if err != nil {
		return nil, err
	}
	return &responseData, nil
}

func (p *PinService) AddPin(cid, name string, origins []interface{}, meta map[string]interface{}) (*RemotePinResponse, error) {
	apiUrl := fmt.Sprintf("%s/pins", p.endpoint)
	return p.PostPin(apiUrl, cid, name, origins, meta)

}

func (p *PinService) ReplacePin(requestId, cid, name string, origins []interface{}, meta map[string]interface{}) (*RemotePinResponse, error) {
	apiUrl := fmt.Sprintf("%s/pins/%s", p.endpoint, requestId)
	return p.PostPin(apiUrl, cid, name, origins, meta)
}

func (p *PinService) GetPin(requestId string) (*RemotePinResponse, error) {
	apiUrl := fmt.Sprintf("%s/pins/%s", p.endpoint, requestId)

	contents, err := p.doApiRequest("GET", apiUrl, nil, pinningSuccessCodes)
	if err != nil {
		return nil, err
	}

	responseData := RemotePinResponse{}
	err = json.Unmarshal(contents, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (p *PinService) RemovePin(requestId string) error {
	apiUrl := fmt.Sprintf("%s/pins/%s", p.endpoint, requestId)
	_, err := p.doApiRequest("DELETE", apiUrl, nil, pinningSuccessCodes)
	return err
}
