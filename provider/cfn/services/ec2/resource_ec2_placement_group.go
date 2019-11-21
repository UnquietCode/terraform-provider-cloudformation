// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2PlacementGroupType string = "AWS::EC2::PlacementGroup"

func ResourceEC2PlacementGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2PlacementGroupRead,
		Create: resourceEC2PlacementGroupCreate,
		Update: resourceEC2PlacementGroupUpdate,
		Delete: resourceEC2PlacementGroupDelete,
		CustomizeDiff: resourceEC2PlacementGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"strategy": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceEC2PlacementGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, meta)
}

func resourceEC2PlacementGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, meta)
}

func resourceEC2PlacementGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, meta)
}

func resourceEC2PlacementGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2PlacementGroupType, data, meta)
}

func resourceEC2PlacementGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2PlacementGroupType, data, meta)
}
