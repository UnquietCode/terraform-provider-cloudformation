// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketNotificationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lambda_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketLambdaConfiguration(),
				Required: false,
			},
			"queue_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketQueueConfiguration(),
				Required: false,
			},
			"topic_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketTopicConfiguration(),
				Required: false,
			},
		},
	}
}