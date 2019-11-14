// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2DHCPOptions() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2DHCPOptionsExists,
		Read: resourceEC2DHCPOptionsRead,
		Create: resourceEC2DHCPOptionsCreate,
		Update: resourceEC2DHCPOptionsUpdate,
		Delete: resourceEC2DHCPOptionsDelete,
		CustomizeDiff: resourceEC2DHCPOptionsCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"netbios_name_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"netbios_node_type": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ntp_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2DHCPOptionsExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2DHCPOptionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::DHCPOptions", ResourceEC2DHCPOptions(), data, meta)
}

func resourceEC2DHCPOptionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::DHCPOptions", ResourceEC2DHCPOptions(), data, meta)
}

func resourceEC2DHCPOptionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::DHCPOptions", ResourceEC2DHCPOptions(), data, meta)
}

func resourceEC2DHCPOptionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::DHCPOptions", data, meta)
}

func resourceEC2DHCPOptionsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
