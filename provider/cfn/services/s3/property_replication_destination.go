// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyReplicationDestination() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_control_translation": {
				Type: schema.TypeList,
				Elem: propertyAccessControlTranslation(),
				Required: false,
				MaxItems: 1,
			},
			"account": {
				Type: schema.TypeString,
				Required: false,
			},
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"encryption_configuration": {
				Type: schema.TypeList,
				Elem: propertyEncryptionConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"storage_class": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}