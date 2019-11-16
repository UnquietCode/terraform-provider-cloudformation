// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var distributionDistributionConfigProperties map[string]string = map[string]string{
	"logging": "Logging",
	"comment": "Comment",
	"default_root_object": "DefaultRootObject",
	"origins": "Origins",
	"viewer_certificate": "ViewerCertificate",
	"price_class": "PriceClass",
	"default_cache_behavior": "DefaultCacheBehavior",
	"custom_error_responses": "CustomErrorResponses",
	"enabled": "Enabled",
	"aliases": "Aliases",
	"ipv6_enabled": "IPV6Enabled",
	"web_acl_id": "WebACLId",
	"http_version": "HttpVersion",
	"restrictions": "Restrictions",
	"cache_behaviors": "CacheBehaviors",
}

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
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyDistributionViewerCertificate(),
				Optional: true,
				MaxItems: 1,
			},
			"price_class": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_cache_behavior": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
