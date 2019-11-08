// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDataCatalogEncryptionSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection_password_encryption": {
				Type: schema.TypeList,
				Elem: propertyConnectionPasswordEncryption(),
				Required: false,
				MaxItems: 1,
			},
			"encryption_at_rest": {
				Type: schema.TypeList,
				Elem: propertyEncryptionAtRest(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}