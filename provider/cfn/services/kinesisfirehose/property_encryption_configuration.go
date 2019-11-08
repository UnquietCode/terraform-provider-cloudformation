// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEncryptionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kms_encryption_config": {
				Type: schema.TypeList,
				Elem: propertyKMSEncryptionConfig(),
				Required: false,
				MaxItems: 1,
			},
			"no_encryption_config": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}