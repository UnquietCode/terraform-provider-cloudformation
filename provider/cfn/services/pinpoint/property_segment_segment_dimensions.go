// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions.html

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySegmentSegmentDimensions(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"demographic": {
				Type: schema.TypeList,
				Elem: propertySegmentDemographic(),
				Optional: true,
				MaxItems: 1,
			},
			"metrics": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"attributes": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"behavior": {
				Type: schema.TypeList,
				Elem: propertySegmentBehavior(),
				Optional: true,
				MaxItems: 1,
			},
			"user_attributes": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"location": {
				Type: schema.TypeList,
				Elem: propertySegmentLocation(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

