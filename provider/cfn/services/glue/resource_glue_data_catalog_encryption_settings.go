// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueDataCatalogEncryptionSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueDataCatalogEncryptionSettingsCreate,
		Read:   resourceGlueDataCatalogEncryptionSettingsRead,
		Update: resourceGlueDataCatalogEncryptionSettingsUpdate,
		Delete: resourceGlueDataCatalogEncryptionSettingsDelete,

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
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueDataCatalogEncryptionSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::DataCatalogEncryptionSettings", ResourceGlueDataCatalogEncryptionSettings(), data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::DataCatalogEncryptionSettings", ResourceGlueDataCatalogEncryptionSettings(), data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::DataCatalogEncryptionSettings", ResourceGlueDataCatalogEncryptionSettings(), data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::DataCatalogEncryptionSettings", data, meta)
}