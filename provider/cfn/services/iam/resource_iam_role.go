// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMRole() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMRoleExists,
		Read:   resourceIAMRoleRead,
		Create: resourceIAMRoleCreate,
		Update: resourceIAMRoleUpdate,
		Delete: resourceIAMRoleDelete,
		
		Schema: map[string]*schema.Schema{
			"assume_role_policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"max_session_duration": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"permissions_boundary": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyRolePolicy(),
				Optional: true,
			},
			"role_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceIAMRoleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::Role", ResourceIAMRole(), data, meta)
}

func resourceIAMRoleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::Role", ResourceIAMRole(), data, meta)
}

func resourceIAMRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::Role", ResourceIAMRole(), data, meta)
}

func resourceIAMRoleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::Role", data, meta)
}