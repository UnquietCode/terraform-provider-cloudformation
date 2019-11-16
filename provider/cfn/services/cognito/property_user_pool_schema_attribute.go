// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolSchemaAttribute(extras...string) *schema.Resource {
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
			"developer_only_attribute": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"mutable": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"attribute_data_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"string_attribute_constraints": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolStringAttributeConstraints(),
				Optional: true,
				MaxItems: 1,
			},
			"required": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"number_attribute_constraints": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolNumberAttributeConstraints(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
