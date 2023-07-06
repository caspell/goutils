package httprequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestAPI struct {
	Host string
}

type ApiError struct {
	URL     string
	Message string
}

func (e *ApiError) Error() string {
	return " Reason: " + e.Message + " , Query: " + e.URL
}

func (api *RestAPI) GetUrl(url string) string {
	return fmt.Sprintf("%s%s", api.Host, url)
}

func NewClient() *http.Client {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    conf,
	}

	return &http.Client{Transport: tr}

}

func (api *RestAPI) NewRequest(method, url string, params map[string]interface{}) ([]byte, error) {

	fullUrl := api.GetUrl(url)

	var buff *bytes.Buffer = nil

	if params != nil {
		if pbytes, err := json.Marshal(params); err != nil {
			return nil, err
		} else {
			buff = bytes.NewBuffer(pbytes)
		}
	}

	req, err := http.NewRequest(method, fullUrl, buff)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Agent")
	req.Header.Add("Content-Type", "application/json")

	client := NewClient()

	resp, err := client.Do(req)

	if err != nil {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("%v \n", closeErr)
		}
		return nil, &ApiError{
			Message: err.Error(),
			URL:     fullUrl,
		}
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (api *RestAPI) Get(url string) (string, error) {

	data, err := api.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	return string(data), nil

}

func (api *RestAPI) GetJson(url string, objType interface{}) (string, error) {

	data, err := api.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(data, &objType); err != nil {
		return "", err
	}

	return string(data), nil
}
