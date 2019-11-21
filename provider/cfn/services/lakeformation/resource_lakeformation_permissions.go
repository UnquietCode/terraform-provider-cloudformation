// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html

package lakeformation

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lakeFormationPermissionsType string = "AWS::LakeFormation::Permissions"

func ResourceLakeFormationPermissions() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceLakeFormationPermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, meta)
}

func resourceLakeFormationPermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, meta)
}

func resourceLakeFormationPermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lakeFormationPermissionsType, ResourceLakeFormationPermissions(), data, meta)
}

func resourceLakeFormationPermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lakeFormationPermissionsType, data, meta)
}

func resourceLakeFormationPermissionsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lakeFormationPermissionsType, data, meta)
}
