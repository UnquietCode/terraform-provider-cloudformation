// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySegmentDemographic() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_version": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"device_type": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"platform": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"channel": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"model": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
			"make": {
				Type: schema.TypeList,
				Elem: propertySegmentSetDimension(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}