// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClassifierCsvClassifier(extras...string) *schema.Resource {
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
			"quote_symbol": {
				Type: schema.TypeString,
				Optional: true,
			},
			"contains_header": {
				Type: schema.TypeString,
				Optional: true,
			},
			"delimiter": {
				Type: schema.TypeString,
				Optional: true,
			},
			"header": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allow_single_column": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"disable_value_trimming": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
