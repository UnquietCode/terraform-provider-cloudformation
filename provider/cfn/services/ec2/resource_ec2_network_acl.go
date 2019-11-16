// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkAclType string = "AWS::EC2::NetworkAcl"

var eC2NetworkAclProperties map[string]string = map[string]string{
	"tags": "Tags",
	"vpc_id": "VpcId",
}

func ResourceEC2NetworkAcl() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2NetworkAclExists,
		Read: resourceEC2NetworkAclRead,
		Create: resourceEC2NetworkAclCreate,
		Update: resourceEC2NetworkAclUpdate,
		Delete: resourceEC2NetworkAclDelete,
		CustomizeDiff: resourceEC2NetworkAclCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
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

func resourceEC2NetworkAclExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2NetworkAclRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkAclType, ResourceEC2NetworkAcl(), data, meta)
}

func resourceEC2NetworkAclCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NetworkAclType, ResourceEC2NetworkAcl(), data, eC2NetworkAclProperties, meta)
}

func resourceEC2NetworkAclUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NetworkAclType, ResourceEC2NetworkAcl(), data, eC2NetworkAclProperties, meta)
}

func resourceEC2NetworkAclDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NetworkAclType, data, meta)
}

func resourceEC2NetworkAclCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NetworkAclType, data, meta)
}
