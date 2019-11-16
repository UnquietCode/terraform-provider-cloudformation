// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html

package lakeformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lakeFormationPermissionsType string = "AWS::LakeFormation::Permissions"

var lakeFormationPermissionsProperties map[string]string = map[string]string{
	"data_lake_principal": "DataLakePrincipal",
	"resource": "Resource",
	"permissions": "Permissions",
	"permissions_with_grant_option": "PermissionsWithGrantOption",
}

func ResourceLakeFormationPermissions() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLakeFormationPermissionsExists,
		Read: resourceLakeFormationPermissionsRead,
		Create: resourceLakeFormationPermissionsCreate,
		Update: resourceLakeFormationPermissionsUpdate,
		Delete: resourceLakeFormationPermissionsDelete,
		CustomizeDiff: resourceLakeFormationPermissionsCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"data_lake_principal": {
				Type: schema.TypeSet,
				Elem: propertyPermissionsDataLakePrincipal(),
				Required: true,
				MaxItems: 1,
			},
			"resource": {
				Type: schema.TypeSet,
				Elem: propertyPermissionsResource(),
				Required: true,
				MaxItems: 1,
			},
			"permissions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"permissions_with_grant_option": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceLakeFormationPermissionsExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLakeFormationPermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, meta)
}

func resourceLakeFormationPermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, lakeFormationPermissionsProperties, meta)
}

func resourceLakeFormationPermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, lakeFormationPermissionsProperties, meta)
}

func resourceLakeFormationPermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lakeFormationPermissionsType, data, meta)
}

func resourceLakeFormationPermissionsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lakeFormationPermissionsType, data, meta)
}
