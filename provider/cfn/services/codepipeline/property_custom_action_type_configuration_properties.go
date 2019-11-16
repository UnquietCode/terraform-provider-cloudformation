// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html

package codepipeline

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var customActionTypeConfigurationPropertiesProperties map[string]string = map[string]string{
	"description": "Description",
	"key": "Key",
	"name": "Name",
	"queryable": "Queryable",
	"required": "Required",
	"secret": "Secret",
	"type": "Type",
}

func propertyCustomActionTypeConfigurationProperties(extras...string) *schema.Resource {
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
			"description": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
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
				Optional: true,
			},
		},
	}
}
