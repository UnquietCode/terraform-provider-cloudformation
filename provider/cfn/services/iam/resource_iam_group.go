// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMGroupType string = "AWS::IAM::Group"

var iAMGroupProperties map[string]string = map[string]string{
	"group_name": "GroupName",
	"managed_policy_arns": "ManagedPolicyArns",
	"path": "Path",
	"policies": "Policies",
}

func ResourceIAMGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMGroupExists,
		Read: resourceIAMGroupRead,
		Create: resourceIAMGroupCreate,
		Update: resourceIAMGroupUpdate,
		Delete: resourceIAMGroupDelete,
		CustomizeDiff: resourceIAMGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeSet,
				Elem: propertyGroupPolicy(),
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

func resourceIAMGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMGroupType, ResourceIAMGroup(), data, meta)
}

func resourceIAMGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMGroupType, ResourceIAMGroup(), data, iAMGroupProperties, meta)
}

func resourceIAMGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMGroupType, ResourceIAMGroup(), data, iAMGroupProperties, meta)
}

func resourceIAMGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMGroupType, data, meta)
}

func resourceIAMGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMGroupType, data, meta)
}
