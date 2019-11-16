// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var distributionCustomOriginConfigProperties map[string]string = map[string]string{
	"origin_read_timeout": "OriginReadTimeout",
	"https_port": "HTTPSPort",
	"origin_keepalive_timeout": "OriginKeepaliveTimeout",
	"origin_ssl_protocols": "OriginSSLProtocols",
	"http_port": "HTTPPort",
	"origin_protocol_policy": "OriginProtocolPolicy",
}

func propertyDistributionCustomOriginConfig(extras...string) *schema.Resource {
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
			"origin_read_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"https_port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"origin_keepalive_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"origin_ssl_protocols": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"http_port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"origin_protocol_policy": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
