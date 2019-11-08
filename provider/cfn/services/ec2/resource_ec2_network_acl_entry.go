// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2NetworkAclEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2NetworkAclEntryCreate,
		Read:   resourceEC2NetworkAclEntryRead,
		Update: resourceEC2NetworkAclEntryUpdate,
		Delete: resourceEC2NetworkAclEntryDelete,

		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"egress": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"icmp": {
				Type: schema.TypeList,
				Elem: propertyIcmp(),
				Required: false,
				MaxItems: 1,
			},
			"ipv6_cidr_block": {
				Type: schema.TypeString,
				Required: false,
			},
			"network_acl_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_range": {
				Type: schema.TypeList,
				Elem: propertyPortRange(),
				Required: false,
				MaxItems: 1,
			},
			"protocol": {
				Type: schema.TypeInt,
				Required: true,
			},
			"rule_action": {
				Type: schema.TypeString,
				Required: true,
			},
			"rule_number": {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NetworkAclEntryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkAclEntry", data, meta)
}

func resourceEC2NetworkAclEntryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkAclEntry", data, meta)
}

func resourceEC2NetworkAclEntryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkAclEntry", data, meta)
}

func resourceEC2NetworkAclEntryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkAclEntry", data, meta)
}