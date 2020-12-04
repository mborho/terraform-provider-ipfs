package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"os"
)

func resourceAdd() *schema.Resource {
	return &schema.Resource{
		Create:        resourceAddCreate,
		Read:          resourceAddRead,
		Delete:        resourceAddDelete,
		CustomizeDiff: resourceAddCustomizeDiff,

		Schema: map[string]*schema.Schema{
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

func resourceAddCustomizeDiff(d *schema.ResourceDiff, m interface{}) error {
	client := m.(*Client)
	id := d.Id()
	filePath := d.Get("path").(string)

	newHash, err := client.getHash(filePath)
	if err != nil {
		return err
	}
	if id != newHash {
		d.SetNewComputed("cid")
	}
	return nil
}

func resourceAddCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("path").(string)

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	cid, err := client.shell.Add(f)
	if err != nil {
		return fmt.Errorf("Error adding object: %s", err)
	}
	d.SetId(cid)
	d.Set("cid", cid)
	return resourceAddRead(d, m)
}

func resourceAddRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("path").(string)
	newHash, err := client.getHash(filePath)
	if err != nil {
		return fmt.Errorf("Error reading object hash: %s", err)
	}
	d.Set("cid", newHash)
	return nil
}

func resourceAddDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
