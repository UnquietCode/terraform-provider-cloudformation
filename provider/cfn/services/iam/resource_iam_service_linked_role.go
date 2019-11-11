// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceIAMServiceLinkedRoleCreate,
		Read:   resourceIAMServiceLinkedRoleRead,
		Update: resourceIAMServiceLinkedRoleUpdate,
		Delete: resourceIAMServiceLinkedRoleDelete,

		Schema: map[string]*schema.Schema{
			"custom_suffix": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"aws_service_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIAMServiceLinkedRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ServiceLinkedRole", ResourceIAMServiceLinkedRole(), data, meta)
}

func resourceIAMServiceLinkedRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ServiceLinkedRole", data, meta)
}