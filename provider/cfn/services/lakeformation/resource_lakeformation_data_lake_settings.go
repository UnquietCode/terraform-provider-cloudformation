// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html

package lakeformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
				Type: schema.TypeList,
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
	return plugin.ResourceRead("AWS::LakeFormation::DataLakeSettings", ResourceLakeFormationDataLakeSettings(), data, meta)
}

func resourceLakeFormationDataLakeSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::LakeFormation::DataLakeSettings", ResourceLakeFormationDataLakeSettings(), data, meta)
}

func resourceLakeFormationDataLakeSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::LakeFormation::DataLakeSettings", ResourceLakeFormationDataLakeSettings(), data, meta)
}

func resourceLakeFormationDataLakeSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::LakeFormation::DataLakeSettings", data, meta)
}

func resourceLakeFormationDataLakeSettingsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::LakeFormation::DataLakeSettings", data, meta)
}

