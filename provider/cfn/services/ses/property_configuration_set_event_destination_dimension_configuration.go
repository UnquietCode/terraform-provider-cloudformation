// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html

package ses

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var configurationSetEventDestinationDimensionConfigurationProperties map[string]string = map[string]string{
	"dimension_value_source": "DimensionValueSource",
	"default_dimension_value": "DefaultDimensionValue",
	"dimension_name": "DimensionName",
}

func propertyConfigurationSetEventDestinationDimensionConfiguration(extras...string) *schema.Resource {
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
			"dimension_value_source": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_dimension_value": {
				Type: schema.TypeString,
				Required: true,
			},
			"dimension_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
