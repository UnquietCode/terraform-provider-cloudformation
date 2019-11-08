// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEventDimensions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics": {
				Type: schema.TypeMap,
				Required: false,
			},
			"event_type": {
				Type: schema.TypeList,
				Elem: propertySetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"attributes": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}