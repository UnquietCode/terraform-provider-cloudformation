// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMGroupExists,
		Read:   resourceIAMGroupRead,
		Create: resourceIAMGroupCreate,
		Update: resourceIAMGroupUpdate,
		Delete: resourceIAMGroupDelete,
		
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
	return plugin.ResourceRead("AWS::IAM::Group", ResourceIAMGroup(), data, meta)
}

func resourceIAMGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::Group", ResourceIAMGroup(), data, meta)
}

func resourceIAMGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::Group", ResourceIAMGroup(), data, meta)
}

func resourceIAMGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::Group", data, meta)
}