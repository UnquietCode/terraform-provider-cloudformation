// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html

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
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compress": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"lambda_function_associations": {
				Type: schema.TypeList,
				Elem: propertyDistributionLambdaFunctionAssociation(),
				Optional: true,
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
				Optional: true,
			},
			"default_ttl": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"field_level_encryption_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"allowed_methods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"cached_methods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"smooth_streaming": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"forwarded_values": {
				Type: schema.TypeList,
				Elem: propertyDistributionForwardedValues(),
				Required: true,
				MaxItems: 1,
			},
			"min_ttl": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"max_ttl": {
				Type: schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

