package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func request(url string) string {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", url, conf)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return ""
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return ""
	}

	return string(buf[:n])
}

func Get(target_url string) (string, error) {

	req, err := http.NewRequest("GET", target_url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "Agent")

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    conf,
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)

	if err != nil {

		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("%v", closeErr)
		}
		return "", err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil

}

func main() {

	url := "https://192.168.0.8:30000/metrics"

	// url = "https://google.com"

	if rst, err := Get(url); err != nil {
		panic(err)
	} else {
		fmt.Println(rst)
	}

}
