// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceLinkedRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceLinkedRoleCreate,
		Read:   resourceServiceLinkedRoleRead,
		Update: resourceServiceLinkedRoleUpdate,
		Delete: resourceServiceLinkedRoleDelete,

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

func resourceServiceLinkedRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceServiceLinkedRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceServiceLinkedRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ServiceLinkedRole", data, meta)
}

func resourceServiceLinkedRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ServiceLinkedRole", data, meta)
}