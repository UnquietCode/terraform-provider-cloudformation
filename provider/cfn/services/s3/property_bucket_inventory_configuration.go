// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html

package s3

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var bucketInventoryConfigurationProperties map[string]string = map[string]string{
	"destination": "Destination",
	"enabled": "Enabled",
	"id": "Id",
	"included_object_versions": "IncludedObjectVersions",
	"optional_fields": "OptionalFields",
	"prefix": "Prefix",
	"schedule_frequency": "ScheduleFrequency",
}

func propertyBucketInventoryConfiguration(extras...string) *schema.Resource {
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
			"destination": {
				Type: schema.TypeList,
				Elem: propertyBucketDestination(),
				Required: true,
				MaxItems: 1,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"included_object_versions": {
				Type: schema.TypeString,
				Required: true,
			},
			"optional_fields": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schedule_frequency": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
