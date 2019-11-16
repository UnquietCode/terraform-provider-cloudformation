// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2PlacementGroupType string = "AWS::EC2::PlacementGroup"

var eC2PlacementGroupProperties map[string]string = map[string]string{
	"strategy": "Strategy",
}

func ResourceEC2PlacementGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2PlacementGroupExists,
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
			},
		},
	}
}

func resourceEC2PlacementGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2PlacementGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, meta)
}

func resourceEC2PlacementGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, eC2PlacementGroupProperties, meta)
}

func resourceEC2PlacementGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2PlacementGroupType, ResourceEC2PlacementGroup(), data, eC2PlacementGroupProperties, meta)
}

func resourceEC2PlacementGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2PlacementGroupType, data, meta)
}

func resourceEC2PlacementGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2PlacementGroupType, data, meta)
}
