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

func propertyQueueConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event": {
				Type: schema.TypeString,
				Required: true,
			},
			"filter": {
				Type: schema.TypeList,
				Elem: propertyNotificationFilter(),
				Required: false,
				MaxItems: 1,
			},
			"queue": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}