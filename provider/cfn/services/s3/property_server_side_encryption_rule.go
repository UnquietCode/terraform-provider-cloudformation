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

func propertyServerSideEncryptionRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_side_encryption_by_default": {
				Type: schema.TypeList,
				Elem: propertyServerSideEncryptionByDefault(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}