package main

import (
	"context"
	"encoding/json"
	"errors"
	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	shell *shell.Shell
	node  string
}

func NewClient(node string) (*Client, error) {
	sh := shell.NewShell(node)
	sh.SetTimeout(10 * time.Minute)
	// return client
	client := &Client{
		shell: sh,
		node:  node,
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
