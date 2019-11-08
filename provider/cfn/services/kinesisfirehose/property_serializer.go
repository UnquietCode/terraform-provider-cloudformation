// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySerializer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"orc_ser_de": {
				Type: schema.TypeList,
				Elem: propertyOrcSerDe(),
				Required: false,
				MaxItems: 1,
			},
			"parquet_ser_de": {
				Type: schema.TypeList,
				Elem: propertyParquetSerDe(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}