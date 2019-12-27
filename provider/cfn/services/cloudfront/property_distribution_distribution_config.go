// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionDistributionConfig(extras...string) *schema.Resource {
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
			"logging": {
				Type: schema.TypeList,
				Elem: propertyDistributionLogging(),
				Optional: true,
				MaxItems: 1,
			},
			"comment": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_root_object": {
				Type: schema.TypeString,
				Optional: true,
			},
			"origins": {
				Type: schema.TypeList,
				Elem: propertyDistributionOrigin(),
				Optional: true,
			},
			"viewer_certificate": {
				Type: schema.TypeList,
				Elem: propertyDistributionViewerCertificate(),
				Optional: true,
				MaxItems: 1,
			},
			"price_class": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_cache_behavior": {
				Type: schema.TypeList,
				Elem: propertyDistributionDefaultCacheBehavior(),
				Optional: true,
				MaxItems: 1,
			},
			"custom_error_responses": {
				Type: schema.TypeList,
				Elem: propertyDistributionCustomErrorResponse(),
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"aliases": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"ipv6_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"web_acl_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"http_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"restrictions": {
				Type: schema.TypeList,
				Elem: propertyDistributionRestrictions(),
				Optional: true,
				MaxItems: 1,
			},
			"cache_behaviors": {
				Type: schema.TypeList,
				Elem: propertyDistributionCacheBehavior(),
				Optional: true,
			},
		},
	}
}
