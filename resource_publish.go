package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourcePublish() *schema.Resource {
	return &schema.Resource{
		Create: resourcePublishCreate,
		Read:   resourcePublishRead,
		Delete: resourcePublishDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "self",
				ForceNew: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePublishCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)
	key := d.Get("key").(string)

	path := fmt.Sprintf("/ipfs/%s", cid)
	resp, err := client.shell.PublishWithDetails(path, key, 0, 0, true)
	if err != nil {
		return err
	}

	d.SetId(resp.Name)
	d.Set("path", path)
	d.Set("name", resp.Name)
	d.Set("value", resp.Value)

	return resourcePublishRead(d, m)
}

func resourcePublishRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	resolvedName, err := client.shell.Resolve(name)
	if err != nil {
		return err
	}

	path := d.Get("path").(string)
	if resolvedName != path {
		d.SetId("")
	}
	return nil
}

func resourcePublishDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
