// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SecurityGroupType string = "AWS::EC2::SecurityGroup"

var eC2SecurityGroupProperties map[string]string = map[string]string{
	"group_description": "GroupDescription",
	"group_name": "GroupName",
	"security_group_egress": "SecurityGroupEgress",
	"security_group_ingress": "SecurityGroupIngress",
	"tags": "Tags",
	"vpc_id": "VpcId",
}

func ResourceEC2SecurityGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2SecurityGroupExists,
		Read: resourceEC2SecurityGroupRead,
		Create: resourceEC2SecurityGroupCreate,
		Update: resourceEC2SecurityGroupUpdate,
		Delete: resourceEC2SecurityGroupDelete,
		CustomizeDiff: resourceEC2SecurityGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_egress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupEgress(),
				Optional: true,
			},
			"security_group_ingress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupIngress(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"vpc_id": {
				Type: schema.TypeString,
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

func resourceEC2SecurityGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2SecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SecurityGroupType, ResourceEC2SecurityGroup(), data, meta)
}

func resourceEC2SecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2SecurityGroupType, ResourceEC2SecurityGroup(), data, eC2SecurityGroupProperties, meta)
}

func resourceEC2SecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2SecurityGroupType, ResourceEC2SecurityGroup(), data, eC2SecurityGroupProperties, meta)
}

func resourceEC2SecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2SecurityGroupType, data, meta)
}

func resourceEC2SecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2SecurityGroupType, data, meta)
}
