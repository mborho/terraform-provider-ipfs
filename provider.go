package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"node": {
				Type:    schema.TypeString,
				Default: "localhost:5001",
				// DefaultFunc: schema.EnvDefaultFunc("IPFS_HOST", nil),
				Optional:    true,
				Description: "ipfs server address, default localhost:5001",
			},

			"remote_pin_service": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},
						"token": {
							Type:     schema.TypeString,
							Required: true,
						},
						"skip_ssl_verify": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ipfs_add":           resourceAdd(),
			"ipfs_dir":           resourceDir(),
			"ipfs_file":          resourceFile(),
			"ipfs_pin":           resourcePin(),
			"ipfs_publish":       resourcePublish(),
			"ipfs_key":           resourceKey(),
			"ipfs_swarm_connect": resourceSwarmConnect(),
			"ipfs_remote_pin":    resourceRemotePin(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	node := d.Get("node").(string)
	remotePinServices := d.Get("remote_pin_service").([]interface{})

	return NewClient(node, remotePinServices)
}
