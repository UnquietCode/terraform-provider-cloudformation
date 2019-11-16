// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html

package lakeformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lakeFormationDataLakeSettingsType string = "AWS::LakeFormation::DataLakeSettings"

var lakeFormationDataLakeSettingsProperties map[string]string = map[string]string{
	"admins": "Admins",
}

func ResourceLakeFormationDataLakeSettings() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLakeFormationDataLakeSettingsExists,
		Read: resourceLakeFormationDataLakeSettingsRead,
		Create: resourceLakeFormationDataLakeSettingsCreate,
		Update: resourceLakeFormationDataLakeSettingsUpdate,
		Delete: resourceLakeFormationDataLakeSettingsDelete,
		CustomizeDiff: resourceLakeFormationDataLakeSettingsCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"admins": {
				Type: schema.TypeSet,
				Elem: propertyDataLakeSettingsAdmins(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLakeFormationDataLakeSettingsExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLakeFormationDataLakeSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lakeFormationDataLakeSettingsType, ResourceLakeFormationDataLakeSettings(), data, meta)
}

func resourceLakeFormationDataLakeSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lakeFormationDataLakeSettingsType, ResourceLakeFormationDataLakeSettings(), data, lakeFormationDataLakeSettingsProperties, meta)
}

func resourceLakeFormationDataLakeSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lakeFormationDataLakeSettingsType, ResourceLakeFormationDataLakeSettings(), data, lakeFormationDataLakeSettingsProperties, meta)
}

func resourceLakeFormationDataLakeSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lakeFormationDataLakeSettingsType, data, meta)
}

func resourceLakeFormationDataLakeSettingsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lakeFormationDataLakeSettingsType, data, meta)
}
