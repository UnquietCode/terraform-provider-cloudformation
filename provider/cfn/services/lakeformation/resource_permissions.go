// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package lakeformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourcePermissionsCreate,
		Read:   resourcePermissionsRead,
		Update: resourcePermissionsUpdate,
		Delete: resourcePermissionsDelete,

		Schema: map[string]*schema.Schema{
			"data_lake_principal": {
				Type: schema.TypeList,
				Elem: propertyPermissionsDataLakePrincipal(),
				Required: true,
				MaxItems: 1,
			},
			"resource": {
				Type: schema.TypeList,
				Elem: propertyPermissionsResource(),
				Required: true,
				MaxItems: 1,
			},
			"permissions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"permissions_with_grant_option": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourcePermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::LakeFormation::Permissions", data, meta)
}

func resourcePermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::LakeFormation::Permissions", data, meta)
}

func resourcePermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::LakeFormation::Permissions", data, meta)
}

func resourcePermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::LakeFormation::Permissions", data, meta)
}