// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-openxjsonserde.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var deliveryStreamOpenXJsonSerDeProperties map[string]string = map[string]string{
	"case_insensitive": "CaseInsensitive",
	"column_to_json_key_mappings": "ColumnToJsonKeyMappings",
	"convert_dots_in_json_keys_to_underscores": "ConvertDotsInJsonKeysToUnderscores",
}

func propertyDeliveryStreamOpenXJsonSerDe(extras...string) *schema.Resource {
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
			"case_insensitive": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"column_to_json_key_mappings": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"convert_dots_in_json_keys_to_underscores": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
