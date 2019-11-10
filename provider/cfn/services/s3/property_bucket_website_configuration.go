// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketWebsiteConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_document": {
				Type: schema.TypeString,
				Required: false,
			},
			"index_document": {
				Type: schema.TypeString,
				Required: false,
			},
			"redirect_all_requests_to": {
				Type: schema.TypeList,
				Elem: propertyBucketRedirectAllRequestsTo(),
				Required: false,
				MaxItems: 1,
			},
			"routing_rules": {
				Type: schema.TypeSet,
				Elem: propertyBucketRoutingRule(),
				Required: false,
			},
		},
	}
}