// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceIAMGroupCreate,
		Read:   resourceIAMGroupRead,
		Update: resourceIAMGroupUpdate,
		Delete: resourceIAMGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"path": {
				Type: schema.TypeString,
				Required: false,
			},
			"policies": {
				Type: schema.TypeSet,
				Elem: propertyGroupPolicy(),
				Required: false,
			},
		},
	}
}

func resourceIAMGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::Group", data, meta)
}

func resourceIAMGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::Group", data, meta)
}

func resourceIAMGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::Group", data, meta)
}

func resourceIAMGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::Group", data, meta)
}