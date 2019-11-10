// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionDefaultCacheBehavior(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compress": {
				Type: schema.TypeBool,
				Required: false,
			},
			"lambda_function_associations": {
				Type: schema.TypeList,
				Elem: propertyDistributionLambdaFunctionAssociation(),
				Required: false,
			},
			"target_origin_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"viewer_protocol_policy": {
				Type: schema.TypeString,
				Required: true,
			},
			"trusted_signers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"default_ttl": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"field_level_encryption_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"allowed_methods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"cached_methods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"smooth_streaming": {
				Type: schema.TypeBool,
				Required: false,
			},
			"forwarded_values": {
				Type: schema.TypeList,
				Elem: propertyDistributionForwardedValues(),
				Required: true,
				MaxItems: 1,
			},
			"min_ttl": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"max_ttl": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}