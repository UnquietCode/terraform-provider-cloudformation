// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html

package wafregional

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyByteMatchSetByteMatchTuple(extras...string) *schema.Resource {
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
			"target_string": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_string_base64": {
				Type: schema.TypeString,
				Optional: true,
			},
			"positional_constraint": {
				Type: schema.TypeString,
				Required: true,
			},
			"text_transformation": {
				Type: schema.TypeString,
				Required: true,
			},
			"field_to_match": {
				Type: schema.TypeList,
				Elem: propertyByteMatchSetFieldToMatch(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}
