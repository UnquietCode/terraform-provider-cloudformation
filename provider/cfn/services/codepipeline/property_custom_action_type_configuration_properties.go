// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html

package codepipeline

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCustomActionTypeConfigurationProperties(extras...string) *schema.Resource {
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
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"key": {
				Type: schema.TypeBool,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"queryable": {
				Type: schema.TypeBool,
				Required: false,
			},
			"required": {
				Type: schema.TypeBool,
				Required: true,
			},
			"secret": {
				Type: schema.TypeBool,
				Required: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}