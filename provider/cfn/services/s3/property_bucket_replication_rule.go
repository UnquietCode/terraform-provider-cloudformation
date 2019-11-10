// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketReplicationRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination": {
				Type: schema.TypeList,
				Elem: propertyBucketReplicationDestination(),
				Required: true,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: false,
			},
			"prefix": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_selection_criteria": {
				Type: schema.TypeList,
				Elem: propertyBucketSourceSelectionCriteria(),
				Required: false,
				MaxItems: 1,
			},
			"status": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}