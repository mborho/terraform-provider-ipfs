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

type PinResponse struct {
	Progress int      `json:"Progress"`
	Pins     []string `json:"Pins"`
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

	var pinResp PinResponse
	req := client.shell.Request("pin/add", cid)
	req.Option("recursive", true)

	if client.temporalToken != "" {
		// X-Hold-Time: 10 - months holdtime
		req.Header("Authorization", fmt.Sprintf("Bearer %s", client.temporalToken))
	}

	err := req.Exec(context.Background(), &pinResp)
	if err != nil {
		return err
	}
	_ = pinResp

	d.SetId(cid)
	return resourcePinRead(d, m)
}

func resourcePinRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)

	if client.temporalToken == "" {
		var resp PinListResponse
		req := client.shell.Request("pin/ls", cid)
		req.Option("quit", true)

		err := req.Exec(context.Background(), &resp)
		if err != nil {
			d.SetId("")
			return nil
		}
		if len(resp.Keys) == 1 && resp.Keys[cid].Type != "indirect" {
			return nil
		}

		d.SetId("")
	}
	return nil
}

func resourcePinDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)

	if client.temporalToken == "" {
		err := client.shell.Unpin(cid)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil
}
