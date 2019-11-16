// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var deliveryStreamParquetSerDeProperties map[string]string = map[string]string{
	"block_size_bytes": "BlockSizeBytes",
	"compression": "Compression",
	"enable_dictionary_compression": "EnableDictionaryCompression",
	"max_padding_bytes": "MaxPaddingBytes",
	"page_size_bytes": "PageSizeBytes",
	"writer_version": "WriterVersion",
}

func propertyDeliveryStreamParquetSerDe(extras...string) *schema.Resource {
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
			"compression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_dictionary_compression": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"max_padding_bytes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"page_size_bytes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"writer_version": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
