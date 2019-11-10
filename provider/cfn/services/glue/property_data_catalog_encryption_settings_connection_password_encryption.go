// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDataCatalogEncryptionSettingsConnectionPasswordEncryption() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"return_connection_password_encrypted": {
				Type: schema.TypeBool,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}