// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2NetworkAclEntry() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2NetworkAclEntryExists,
		Read:   resourceEC2NetworkAclEntryRead,
		Create: resourceEC2NetworkAclEntryCreate,
		Update: resourceEC2NetworkAclEntryUpdate,
		Delete: resourceEC2NetworkAclEntryDelete,
		
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"egress": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"icmp": {
				Type: schema.TypeList,
				Elem: propertyNetworkAclEntryIcmp(),
				Optional: true,
				MaxItems: 1,
			},
			"ipv6_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_acl_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"port_range": {
				Type: schema.TypeList,
				Elem: propertyNetworkAclEntryPortRange(),
				Optional: true,
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NetworkAclEntryExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2NetworkAclEntryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkAclEntry", ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkAclEntry", ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkAclEntry", ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkAclEntry", data, meta)
}