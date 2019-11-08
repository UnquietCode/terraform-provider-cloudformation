// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMRoleCreate,
		Read:   resourceIAMRoleRead,
		Update: resourceIAMRoleUpdate,
		Delete: resourceIAMRoleDelete,

		Schema: map[string]*schema.Schema{
			"assume_role_policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"max_session_duration": {
				Type: schema.TypeInt,
				Required: false,
			},
			"path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"permissions_boundary": {
				Type: schema.TypeString,
				Required: false,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyPolicy(),
				Required: false,
			},
			"role_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceIAMRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::Role", data, meta)
}

func resourceIAMRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::Role", data, meta)
}

func resourceIAMRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::Role", data, meta)
}

func resourceIAMRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::Role", data, meta)
}