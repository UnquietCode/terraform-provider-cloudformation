// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceEC2NetworkAclEntry() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2NetworkAclEntryRead,
		
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type: schema.TypeString,
				Optional: true,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2NetworkAclEntryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkAclEntryType, DatasourceEC2NetworkAclEntry(), data, meta)
}
