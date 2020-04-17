package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	shell         *shell.Shell
	temporalToken string
	node          string
}

func NewTemporalToken(creds map[string]interface{}) (string, error) { //*http.Client, error) {
	url := "https://api.temporal.cloud/v2/auth/login"

	reqBody, err := json.Marshal(map[string]string{
		"username": fmt.Sprintf("%s", creds["username"]),
		"password": fmt.Sprintf("%s", creds["password"]),
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Authentication to Temporal failed!")
	}

	var data map[string]string
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data["token"], nil
}

func NewClient(node string, temporal map[string]interface{}) (*Client, error) {
	sh := shell.NewShell(node)
	sh.SetTimeout(10 * time.Minute)

	client := &Client{
		shell: sh,
		node:  node,
	}

	// add temporal token token if neccessary
	if len(temporal) > 0 {
		temporalToken, err := NewTemporalToken(temporal)
		if err != nil {
			return nil, err
		}
		client.temporalToken = temporalToken
	}
	return client, nil
}

func (c *Client) getHash(path string) (string, error) {
	f, err := os.Open(path) //, os.O_RDWR, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return c.shell.Add(f, shell.OnlyHash(true))
}

func (c *Client) getHashDir(dir string) (string, error) {
	stat, err := os.Lstat(dir)
	if err != nil {
		return "", err
	}

	sf, err := files.NewSerialFile(dir, false, stat)
	if err != nil {
		return "", err
	}
	slf := files.NewSliceDirectory([]files.DirEntry{files.FileEntry(filepath.Base(dir), sf)})
	reader := files.NewMultiFileReader(slf, true)

	resp, err := c.shell.Request("add").
		Option("recursive", true).
		Option("hash-only", true).
		Body(reader).
		Send(context.Background())
	if err != nil {
		return "", nil
	}

	defer resp.Close()
	log.Println("OUT:", resp.Output)

	if resp.Error != nil {
		return "", resp.Error
	}

	dec := json.NewDecoder(resp.Output)
	var final string
	for {
		var out object
		err = dec.Decode(&out)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		final = out.Hash
	}

	if final == "" {
		return "", errors.New("no results received")
	}

	return final, nil
	//return "", nil //c.shell.Add(f, shell.OnlyHash(true), Recursive(true))
}

type object struct {
	Hash string
}

func Recursive(enabled bool) shell.AddOpts {
	return func(rb *shell.RequestBuilder) error {
		rb.Option("recursive", enabled)
		return nil
	}
}
