package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceDir() *schema.Resource {
	return &schema.Resource{
		Create:        resourceDirCreate,
		Read:          resourceDirRead,
		Delete:        resourceDirDelete,
		CustomizeDiff: resourceDirCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"local_path": &schema.Schema{
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

func resourceDirCustomizeDiff(d *schema.ResourceDiff, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("local_path").(string)
	id := d.Id()

	newHash, err := client.getHashDir(filePath)
	if err != nil {
		return err
	}
	if id != newHash {
		log.Println("######## DO UPDATE ###########")
		log.Printf("oldHash %s\n", id)
		log.Printf("newHash %s\n", newHash)
		d.SetNewComputed("cid")
	}
	return nil
}

func resourceDirCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	dirPath := d.Get("local_path").(string)
	cid, err := client.shell.AddDir(dirPath)
	if err != nil {
		return fmt.Errorf("Error adding directory: %s", err)
	}
	d.SetId(cid)
	d.Set("cid", cid)
	return resourceDirRead(d, m)
}

func resourceDirRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("local_path").(string)

	newHash, err := client.getHashDir(filePath)
	if err != nil {
		return fmt.Errorf("Error reading directory hash: %s", err)
	}
	d.Set("cid", newHash)
	return nil
}

func resourceDirDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Id()
	d.SetId("")
	err := client.shell.Unpin(cid)
	if err != nil {
		log.Println("Error unpinning directory: ", err)
	}
	return nil
}
