package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyCreate,
		Read:   resourceKeyRead,
		Delete: resourceKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "rsa",
				ForceNew: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt, //Float,
				Optional: true,
				Default:  2048,
				ForceNew: true,
			},
		},
	}
}

type KeyResponse struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type KeyListResponse struct {
	Keys []KeyResponse
}

func resourceKeyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	name := d.Get("name").(string)
	keyType := d.Get("type").(string)
	size := d.Get("size").(int)

	var keyResp KeyResponse
	req := client.shell.Request("key/gen", name)
	req.Option("type", keyType)
	req.Option("size", size)

	err := req.Exec(context.Background(), &keyResp)
	if err != nil {
		return err
	}

	d.SetId(keyResp.Id)
	return resourceKeyRead(d, m)
}

func resourceKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	name := d.Get("name").(string)

	var resp KeyListResponse
	req := client.shell.Request("key/list")

	err := req.Exec(context.Background(), &resp)
	if err != nil {
		return err
	}
	for _, k := range resp.Keys {
		if k.Name == name {
			// key exists
			return nil
		}
	}
	d.SetId("")
	return nil
}

func resourceKeyDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	name := d.Get("name").(string)

	var resp KeyListResponse
	req := client.shell.Request("key/rm", name)

	err := req.Exec(context.Background(), &resp)
	if err != nil {
		return err
	}
	for _, k := range resp.Keys {
		if k.Name == name {
			// key deleted
			d.SetId("")
			return nil
		}
	}
	return fmt.Errorf("Key was not deleted:", err)
}
