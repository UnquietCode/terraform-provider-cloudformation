// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGroupGroupVersion(extras...string) *schema.Resource {
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
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}