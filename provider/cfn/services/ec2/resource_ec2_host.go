// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2Host() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2HostCreate,
		Read:   resourceEC2HostRead,
		Update: resourceEC2HostUpdate,
		Delete: resourceEC2HostDelete,

		Schema: map[string]*schema.Schema{
			"auto_placement": {
				Type: schema.TypeString,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"host_recovery": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2HostCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Host", data, meta)
}

func resourceEC2HostRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Host", data, meta)
}

func resourceEC2HostUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Host", data, meta)
}

func resourceEC2HostDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Host", data, meta)
}