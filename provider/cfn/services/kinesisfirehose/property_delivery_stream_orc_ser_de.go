// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamOrcSerDe(extras...string) *schema.Resource {
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
			"block_size_bytes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"bloom_filter_columns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"bloom_filter_false_positive_probability": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"compression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"dictionary_key_threshold": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"enable_padding": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"format_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"padding_tolerance": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"row_index_stride": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"stripe_size_bytes": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}

