// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func propertyApplicationPropertyGroup(extras...string) *schema.Resource {
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
			"property_map": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"property_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
