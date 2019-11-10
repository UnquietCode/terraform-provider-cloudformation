// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
		},
	}
}

func resourceGlueDataCatalogEncryptionSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::DataCatalogEncryptionSettings", data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::DataCatalogEncryptionSettings", data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::DataCatalogEncryptionSettings", data, meta)
}

func resourceGlueDataCatalogEncryptionSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::DataCatalogEncryptionSettings", data, meta)
}