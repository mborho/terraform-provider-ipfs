package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSwarmConnect() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwarmConnectCreate,
		Read:   resourceSwarmConnectRead,
		Delete: resourceSwarmConnectDelete,

		Schema: map[string]*schema.Schema{
			"addresses": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}
}

func resourceSwarmConnectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	addresses := d.Get("addresses").([]interface{})
	for _, addr := range addresses {
		err := client.shell.SwarmConnect(context.Background(), addr.(string))
		if err != nil {
			return err
		}

	}
	id := hashcode.String(fmt.Sprintf("%s", addresses))
	d.SetId(string(id))
	return resourceSwarmConnectRead(d, m)
}

func resourceSwarmConnectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSwarmConnectDelete(d *schema.ResourceData, m interface{}) error {
	// Connections will get garbage collected by IPFS node
	d.SetId("")
	return nil
}
