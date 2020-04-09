package main

import (
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	//"os"
	//"path/filepath"
)

func resourceDir() *schema.Resource {
	return &schema.Resource{
		Create: resourceDirCreate,
		Read:   resourceDirRead,
		//Update:        resourceDirUpdate,
		Delete:        resourceDirDelete,
		CustomizeDiff: resourceDirCustomizeDiff,

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

func resourceDirCustomizeDiff(d *schema.ResourceDiff, m interface{}) error {
	client := m.(*Client)
	filePath := d.Get("path").(string)
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
	dirPath := d.Get("path").(string)
	cid, err := client.shell.AddDir(dirPath)
	if err != nil {
		return err
	}
	d.SetId(cid)
	d.Set("cid", cid)
	return resourceDirRead(d, m)
}

func resourceDirRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	//cid := d.Id()
	filePath := d.Get("path").(string)

	newHash, err := client.getHashDir(filePath)
	if err != nil {
		return err
	}
	d.Set("cid", newHash)
	return nil
}

func resourceDirDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("######## DELETE ###########")
	client := m.(*Client)
	cid := d.Id()
	d.SetId("")
	err := client.shell.Unpin(cid)
	if err != nil {
		log.Println("Error deleting ", err)
		//return err
	}
	return nil
}
