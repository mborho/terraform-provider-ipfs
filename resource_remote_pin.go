package main

import (
	//"context"
	//"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "log"
)

func resourceRemotePin() *schema.Resource {
	return &schema.Resource{
		Create: resourceRemotePinCreate,
		Read:   resourceRemotePinRead,
		Update: resourceRemotePinUpdate,
		Delete: resourceRemotePinDelete,

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"origins": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"meta": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			/*"info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},*/
			"delegates": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceRemotePinCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	cid := d.Get("cid").(string)
	name := d.Get("name").(string)
	origins := d.Get("origins").([]interface{})
	meta := d.Get("meta").(map[string]interface{})
	service := d.Get("service").(string)

	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("Pin service %s unknown!", service)
	}

	resp, err := pinClient.AddPin(cid, name, origins, meta)
	if err != nil {
		return fmt.Errorf("Error when calling remote pin service: %s", err)
	}

	d.SetId(resp.RequestId)
	d.Set("request_id", resp.RequestId)
	d.Set("status", resp.Status)
	d.Set("delegates", resp.Delegates)
	return resourceRemotePinRead(d, m)
}

func resourceRemotePinRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	requestId := d.Get("request_id").(string)
	service := d.Get("service").(string)

	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("Pin service %s unknown!", service)
	}

	resp, err := pinClient.GetPin(requestId)
	if err != nil {
		return fmt.Errorf("Error when calling remote pin service: %s", err)
	}

	d.Set("request_id", resp.RequestId)
	d.Set("status", resp.Status)
	d.Set("delegates", resp.Delegates)
	return nil
}

func resourceRemotePinUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	requestId := d.Get("request_id").(string)
	service := d.Get("service").(string)
	cid := d.Get("cid").(string)
	name := d.Get("name").(string)
	origins := d.Get("origins").([]interface{})
	meta := d.Get("meta").(map[string]interface{})

	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("Pin service %s unknown!", service)
	}

	resp, err := pinClient.ReplacePin(requestId, cid, name, origins, meta)
	if err != nil {
		return fmt.Errorf("Error when calling remote pin service: %s", err)
	}

	d.SetId(resp.RequestId)
	d.Set("request_id", resp.RequestId)
	d.Set("status", resp.Status)
	d.Set("delegates", resp.Delegates)
	return nil
}

func resourceRemotePinDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	requestId := d.Get("request_id").(string)
	service := d.Get("service").(string)
	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("Pin service %s unknown!", service)
	}

	err := pinClient.RemovePin(requestId)
	if err != nil {
		return fmt.Errorf("Error when calling remote pin service: %s", err)
	}

	d.SetId("")
	return nil
}
