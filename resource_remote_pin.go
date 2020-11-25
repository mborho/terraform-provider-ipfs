package main

import (
	//"context"
	//"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceRemotePin() *schema.Resource {
	return &schema.Resource{
		Create: resourceRemotePinCreate,
		Read:   resourceRemotePinRead,
		//Update: resourceRemotePinUpdate,
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
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"origins": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			/*"meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},*/
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
	service := d.Get("service").(string)

	log.Println("#########################\n####################\n#########")
	log.Println("CID: ", cid)
	log.Println("Name: ", name)
	log.Println("Service: ", service)
	log.Println("Origins: ", origins)
	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("load client for pin service %s failed!", service)
	}
	resp, err := pinClient.AddPin(cid, name, origins)
	if err != nil {
		return nil
	}
	log.Println("DATA %+v", resp)
	d.SetId(resp.RequestId)
	d.Set("request_id", resp.RequestId)
	d.Set("status", resp.Status)
	return resourceRemotePinRead(d, m)
}

func resourceRemotePinRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

/*func resourceRemotePinUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}*/

func resourceRemotePinDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	requestId := d.Get("request_id").(string)
	service := d.Get("service").(string)
	pinClient, ok := client.pinServices[service]
	if ok != true {
		return fmt.Errorf("load client for pin service %s failed!", service)
	}

	err := pinClient.RemovePin(requestId)
	if err != nil {
		return nil
	}

	d.SetId("")
	return nil
}
