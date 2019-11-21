// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html

package dax

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dAXSubnetGroupType string = "AWS::DAX::SubnetGroup"

func ResourceDAXSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceDAXSubnetGroupRead,
		Create: resourceDAXSubnetGroupCreate,
		Update: resourceDAXSubnetGroupUpdate,
		Delete: resourceDAXSubnetGroupDelete,
		CustomizeDiff: resourceDAXSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceDAXSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dAXSubnetGroupType, ResourceDAXSubnetGroup(), data, meta)
}

func resourceDAXSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dAXSubnetGroupType, ResourceDAXSubnetGroup(), data, meta)
}

func resourceDAXSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dAXSubnetGroupType, ResourceDAXSubnetGroup(), data, meta)
}

func resourceDAXSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dAXSubnetGroupType, data, meta)
}

func resourceDAXSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dAXSubnetGroupType, data, meta)
}
