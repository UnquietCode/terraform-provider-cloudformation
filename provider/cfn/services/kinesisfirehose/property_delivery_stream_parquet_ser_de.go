// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamParquetSerDe() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_size_bytes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"compression": {
				Type: schema.TypeString,
				Required: false,
			},
			"enable_dictionary_compression": {
				Type: schema.TypeBool,
				Required: false,
			},
			"max_padding_bytes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"page_size_bytes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"writer_version": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}