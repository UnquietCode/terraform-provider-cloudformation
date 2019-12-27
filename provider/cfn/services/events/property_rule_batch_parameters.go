// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html

package events

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRuleBatchParameters(extras...string) *schema.Resource {
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
			"array_properties": {
				Type: schema.TypeList,
				Elem: propertyRuleBatchArrayProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"job_definition": {
				Type: schema.TypeString,
				Required: true,
			},
			"job_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"retry_strategy": {
				Type: schema.TypeList,
				Elem: propertyRuleBatchRetryStrategy(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
