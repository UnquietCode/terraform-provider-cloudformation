// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html

package mediastore

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var containerCorsRuleProperties map[string]string = map[string]string{
	"allowed_methods": "AllowedMethods",
	"allowed_origins": "AllowedOrigins",
	"expose_headers": "ExposeHeaders",
	"max_age_seconds": "MaxAgeSeconds",
	"allowed_headers": "AllowedHeaders",
}

func propertyContainerCorsRule(extras...string) *schema.Resource {
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
			"allowed_methods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allowed_origins": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"expose_headers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"max_age_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"allowed_headers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}
