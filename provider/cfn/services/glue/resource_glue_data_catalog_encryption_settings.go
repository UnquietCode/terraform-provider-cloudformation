// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueDataCatalogEncryptionSettingsType string = "AWS::Glue::DataCatalogEncryptionSettings"

var glueDataCatalogEncryptionSettingsProperties map[string]string = map[string]string{
	"data_catalog_encryption_settings": "DataCatalogEncryptionSettings",
	"catalog_id": "CatalogId",
}

func ResourceGlueDataCatalogEncryptionSettings() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueDataCatalogEncryptionSettingsExists,
		Read: resourceGlueDataCatalogEncryptionSettingsRead,
		Create: resourceGlueDataCatalogEncryptionSettingsCreate,
		Update: resourceGlueDataCatalogEncryptionSettingsUpdate,
		Delete: resourceGlueDataCatalogEncryptionSettingsDelete,
		CustomizeDiff: resourceGlueDataCatalogEncryptionSettingsCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"data_catalog_encryption_settings": {
				Type: schema.TypeList,
				Elem: propertyDataCatalogEncryptionSettingsDataCatalogEncryptionSettings(),
				Required: true,
				MaxItems: 1,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueDataCatalogEncryptionSettingsExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueDataCatalogEncryptionSettingsType, ResourceGlueDataCatalogEncryptionSettings(), data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueDataCatalogEncryptionSettingsType, ResourceGlueDataCatalogEncryptionSettings(), data, glueDataCatalogEncryptionSettingsProperties, meta)
}

func resourceGlueDataCatalogEncryptionSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueDataCatalogEncryptionSettingsType, ResourceGlueDataCatalogEncryptionSettings(), data, glueDataCatalogEncryptionSettingsProperties, meta)
}

func resourceGlueDataCatalogEncryptionSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueDataCatalogEncryptionSettingsType, data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueDataCatalogEncryptionSettingsType, data, meta)
}
