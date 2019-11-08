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

func propertySegmentDimensions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"demographic": {
				Type: schema.TypeList,
				Elem: propertyDemographic(),
				Required: false,
				MaxItems: 1,
			},
			"metrics": {
				Type: schema.TypeMap,
				Required: false,
			},
			"attributes": {
				Type: schema.TypeMap,
				Required: false,
			},
			"behavior": {
				Type: schema.TypeList,
				Elem: propertyBehavior(),
				Required: false,
				MaxItems: 1,
			},
			"user_attributes": {
				Type: schema.TypeMap,
				Required: false,
			},
			"location": {
				Type: schema.TypeList,
				Elem: propertyLocation(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}