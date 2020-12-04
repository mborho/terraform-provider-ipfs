package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type PinInfo struct {
	Type string `json:"Type"`
}

type PinListResponse struct {
	Keys map[string]PinInfo `json:"Keys"`
}

func resourcePin() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinCreate,
		Read:   resourcePinRead,
		Delete: resourcePinDelete,

		Schema: map[string]*schema.Schema{
			"cid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)

	err := client.shell.Pin(cid)
	if err != nil {
		return fmt.Errorf("Error creating pin: %s", err)
	}
	d.SetId(cid)
	return resourcePinRead(d, m)
}

func resourcePinRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)

	var resp PinListResponse
	req := client.shell.Request("pin/ls", cid)
	req.Option("quit", true)

	err := req.Exec(context.Background(), &resp)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading pin: %s", err)
	}
	if len(resp.Keys) == 1 && resp.Keys[cid].Type != "indirect" {
		return nil
	}

	d.SetId("")
	return nil
}

func resourcePinDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)

	err := client.shell.Unpin(cid)
	if err != nil {
		return fmt.Errorf("Error deleting pin: %s", err)
	}

	d.SetId("")
	return nil
}
