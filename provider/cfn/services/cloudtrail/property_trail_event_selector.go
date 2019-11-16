// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html

package cloudtrail

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var trailEventSelectorProperties map[string]string = map[string]string{
	"data_resources": "DataResources",
	"include_management_events": "IncludeManagementEvents",
	"read_write_type": "ReadWriteType",
}

func propertyTrailEventSelector(extras...string) *schema.Resource {
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
			"data_resources": {
				Type: schema.TypeSet,
				Elem: propertyTrailDataResource(),
				Optional: true,
			},
			"include_management_events": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"read_write_type": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
