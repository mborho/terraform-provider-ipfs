package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRemotePin() *schema.Resource {
	return &schema.Resource{
		Create: resourceRemotePinCreate,
		Read:   resourceRemotePinRead,
		Delete: resourceRemotePinDelete,

		Schema: map[string]*schema.Schema{
			"cid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"origins": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"info": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"delegates": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceRemotePinCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("requestId")
	return resourceRemotePinRead(d, m)
}

func resourceRemotePinRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRemotePinDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
