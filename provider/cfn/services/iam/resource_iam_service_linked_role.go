// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMServiceLinkedRole() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMServiceLinkedRoleExists,
		Read: resourceIAMServiceLinkedRoleRead,
		Create: resourceIAMServiceLinkedRoleCreate,
		Update: resourceIAMServiceLinkedRoleUpdate,
		Delete: resourceIAMServiceLinkedRoleDelete,
		CustomizeDiff: resourceIAMServiceLinkedRoleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"custom_suffix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"aws_service_name": {
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

func resourceIAMServiceLinkedRoleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMServiceLinkedRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceIAMServiceLinkedRoleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

