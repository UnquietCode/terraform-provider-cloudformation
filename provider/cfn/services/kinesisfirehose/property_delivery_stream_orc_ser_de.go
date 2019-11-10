// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_size_bytes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"bloom_filter_columns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"bloom_filter_false_positive_probability": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"compression": {
				Type: schema.TypeString,
				Required: false,
			},
			"dictionary_key_threshold": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"enable_padding": {
				Type: schema.TypeBool,
				Required: false,
			},
			"format_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"padding_tolerance": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"row_index_stride": {
				Type: schema.TypeInt,
				Required: false,
			},
			"stripe_size_bytes": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}