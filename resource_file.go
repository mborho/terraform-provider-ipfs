package main

import (
	"context"
	_ "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"os"

	files "github.com/ipfs/go-ipfs-files"
)

type FileResponse struct {
	Hash string `json:"Hash"`
	Name string `json:"Name"`
	Size int64  `json:"Size"`
	Type int    `json:"Type"`
}

type FileListResponse struct {
	Entries []FileResponse
}

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("file").(string)
	path := d.Get("path").(string)

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	fr := files.NewReaderFile(f)
	slf := files.NewSliceDirectory([]files.DirEntry{files.FileEntry("", fr)})
	fileReader := files.NewMultiFileReader(slf, true)

	req := client.shell.Request("files/write", path)
	req.Body(fileReader)
	req.Option("create", true)
	req.Option("parents", true)

	err = req.Exec(context.Background(), nil)
	if err != nil {
		return err
	}
	d.SetId(path)
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	path := d.Get("path").(string)

	var resp FileListResponse
	req := client.shell.Request("files/ls", path)
	req.Option("l", true)
	if err := req.Exec(context.Background(), &resp); err != nil {
		return err
	}

	for _, e := range resp.Entries {
		log.Printf("ENTRY: %+v %s\n", e, path)
		d.Set("cid", e.Hash)
		// file exists
		return nil
	}
	d.SetId("")
	return nil
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	path := d.Get("path").(string)

	req := client.shell.Request("files/rm", path)
	if err := req.Exec(context.Background(), nil); err != nil {
		return err
	}

	d.SetId("")
	return nil
}
