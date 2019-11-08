// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging": {
				Type: schema.TypeList,
				Elem: propertyLogging(),
				Required: false,
				MaxItems: 1,
			},
			"comment": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_root_object": {
				Type: schema.TypeString,
				Required: false,
			},
			"origins": {
				Type: schema.TypeList,
				Elem: propertyOrigin(),
				Required: false,
			},
			"viewer_certificate": {
				Type: schema.TypeList,
				Elem: propertyViewerCertificate(),
				Required: false,
				MaxItems: 1,
			},
			"price_class": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_cache_behavior": {
				Type: schema.TypeList,
				Elem: propertyDefaultCacheBehavior(),
				Required: false,
				MaxItems: 1,
			},
			"custom_error_responses": {
				Type: schema.TypeList,
				Elem: propertyCustomErrorResponse(),
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"aliases": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"ipv6_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"web_acl_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"http_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"restrictions": {
				Type: schema.TypeList,
				Elem: propertyRestrictions(),
				Required: false,
				MaxItems: 1,
			},
			"cache_behaviors": {
				Type: schema.TypeList,
				Elem: propertyCacheBehavior(),
				Required: false,
			},
		},
	}
}