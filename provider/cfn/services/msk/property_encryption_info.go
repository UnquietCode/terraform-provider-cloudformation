// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package msk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEncryptionInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encryption_at_rest": {
				Type: schema.TypeList,
				Elem: propertyEncryptionAtRest(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"encryption_in_transit": {
				Type: schema.TypeList,
				Elem: propertyEncryptionInTransit(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}