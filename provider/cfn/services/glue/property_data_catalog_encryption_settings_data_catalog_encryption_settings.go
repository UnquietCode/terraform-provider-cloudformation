// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var dataCatalogEncryptionSettingsDataCatalogEncryptionSettingsProperties map[string]string = map[string]string{
	"connection_password_encryption": "ConnectionPasswordEncryption",
	"encryption_at_rest": "EncryptionAtRest",
}

func propertyDataCatalogEncryptionSettingsDataCatalogEncryptionSettings(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection_password_encryption": {
				Type: schema.TypeSet,
				Elem: propertyDataCatalogEncryptionSettingsConnectionPasswordEncryption(),
				Optional: true,
				MaxItems: 1,
			},
			"encryption_at_rest": {
				Type: schema.TypeSet,
				Elem: propertyDataCatalogEncryptionSettingsEncryptionAtRest(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
