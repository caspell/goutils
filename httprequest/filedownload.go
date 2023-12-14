package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

type DownloadFile struct {
	FileName string
	FilePath string
	FileUrl  string
	Size     int64
	File     *os.File
}

func (f *DownloadFile) Download() error {

	fileURL, err := url.Parse(f.FileUrl)
	if err != nil {
		return err
	}
	path := fileURL.Path
	paths := strings.Split(path, "/")
	f.FileName = paths[len(paths)-1]

	if f.File, err = os.Create(f.FileName); err != nil {
		return err
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(f.FileUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	size, err := io.Copy(f.File, resp.Body)
	defer f.File.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
