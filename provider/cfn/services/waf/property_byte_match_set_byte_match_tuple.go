// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html

package waf

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var byteMatchSetByteMatchTupleProperties map[string]string = map[string]string{
	"field_to_match": "FieldToMatch",
	"positional_constraint": "PositionalConstraint",
	"target_string": "TargetString",
	"target_string_base64": "TargetStringBase64",
	"text_transformation": "TextTransformation",
}

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
			"field_to_match": {
				Type: schema.TypeSet,
				Elem: propertyByteMatchSetFieldToMatch(),
				Required: true,
				MaxItems: 1,
			},
			"positional_constraint": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_string": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_string_base64": {
				Type: schema.TypeString,
				Optional: true,
			},
			"text_transformation": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
