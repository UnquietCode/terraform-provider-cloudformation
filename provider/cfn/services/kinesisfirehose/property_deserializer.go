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

func propertyDeserializer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hive_json_ser_de": {
				Type: schema.TypeList,
				Elem: propertyHiveJsonSerDe(),
				Required: false,
				MaxItems: 1,
			},
			"open_x_json_ser_de": {
				Type: schema.TypeList,
				Elem: propertyOpenXJsonSerDe(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}