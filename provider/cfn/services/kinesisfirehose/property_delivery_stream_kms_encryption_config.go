// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamKMSEncryptionConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"awskms_key_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}