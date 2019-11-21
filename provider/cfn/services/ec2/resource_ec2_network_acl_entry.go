// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkAclEntryType string = "AWS::EC2::NetworkAclEntry"

func ResourceEC2NetworkAclEntry() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2NetworkAclEntryRead,
		Create: resourceEC2NetworkAclEntryCreate,
		Update: resourceEC2NetworkAclEntryUpdate,
		Delete: resourceEC2NetworkAclEntryDelete,
		CustomizeDiff: resourceEC2NetworkAclEntryCustomizeDiff,
		
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
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2NetworkAclEntryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkAclEntryType, ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NetworkAclEntryType, ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NetworkAclEntryType, ResourceEC2NetworkAclEntry(), data, meta)
}

func resourceEC2NetworkAclEntryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NetworkAclEntryType, data, meta)
}

func resourceEC2NetworkAclEntryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NetworkAclEntryType, data, meta)
}
