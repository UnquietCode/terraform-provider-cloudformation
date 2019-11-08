// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"aws_service_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIAMServiceLinkedRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceIAMServiceLinkedRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceIAMServiceLinkedRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceIAMServiceLinkedRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ServiceLinkedRole", data, meta)
}