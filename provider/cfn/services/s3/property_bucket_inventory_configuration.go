// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketInventoryConfiguration() *schema.Resource {
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
				Required: false,
				Set: schema.HashString,
			},
			"prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"schedule_frequency": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}