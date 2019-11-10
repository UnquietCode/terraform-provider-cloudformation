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

func ResourceDataLakeSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataLakeSettingsCreate,
		Read:   resourceDataLakeSettingsRead,
		Update: resourceDataLakeSettingsUpdate,
		Delete: resourceDataLakeSettingsDelete,

		Schema: map[string]*schema.Schema{
			"admins": {
				Type: schema.TypeList,
				Elem: propertyDataLakeSettingsAdmins(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceDataLakeSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::LakeFormation::DataLakeSettings", data, meta)
}

func resourceDataLakeSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::LakeFormation::DataLakeSettings", data, meta)
}

func resourceDataLakeSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::LakeFormation::DataLakeSettings", data, meta)
}

func resourceDataLakeSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::LakeFormation::DataLakeSettings", data, meta)
}