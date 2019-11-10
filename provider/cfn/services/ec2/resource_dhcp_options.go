// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDHCPOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPOptionsCreate,
		Read:   resourceDHCPOptionsRead,
		Update: resourceDHCPOptionsUpdate,
		Delete: resourceDHCPOptionsDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"domain_name_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"netbios_name_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"netbios_node_type": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"ntp_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDHCPOptionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::DHCPOptions", data, meta)
}

func resourceDHCPOptionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::DHCPOptions", data, meta)
}

func resourceDHCPOptionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::DHCPOptions", data, meta)
}

func resourceDHCPOptionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::DHCPOptions", data, meta)
}