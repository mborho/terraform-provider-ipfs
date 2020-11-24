package main

import (
	//"context"
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
			"meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
	if pinClient, ok := client.pinServices[service]; ok {
		log.Println("PinClient: ", pinClient)
		//do something here
	} else {
		return fmt.Errorf("load client for pin service %s failed!", service)
	}
	log.Println("#########################\n####################\n#########")
	_ = client
	d.SetId("requestId")

	//var dels []interface{}
	/*dels := make([]interface{}, 0)

	dels = append(dels, "one")
	dels = append(dels, "two")
	s := []interface{}{}
	s = append(s, "one")
	d.Set("delegates", s)*/
	return resourceRemotePinRead(d, m)
}

func resourceRemotePinRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

/*func resourceRemotePinUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}*/

func resourceRemotePinDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
