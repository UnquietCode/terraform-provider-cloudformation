// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStorageDescriptor() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"stored_as_sub_directories": {
				Type: schema.TypeBool,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"bucket_columns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"skewed_info": {
				Type: schema.TypeList,
				Elem: propertySkewedInfo(),
				Required: false,
				MaxItems: 1,
			},
			"input_format": {
				Type: schema.TypeString,
				Required: false,
			},
			"number_of_buckets": {
				Type: schema.TypeInt,
				Required: false,
			},
			"output_format": {
				Type: schema.TypeString,
				Required: false,
			},
			"columns": {
				Type: schema.TypeList,
				Elem: propertyColumn(),
				Required: false,
			},
			"serde_info": {
				Type: schema.TypeList,
				Elem: propertySerdeInfo(),
				Required: false,
				MaxItems: 1,
			},
			"sort_columns": {
				Type: schema.TypeList,
				Elem: propertyOrder(),
				Required: false,
			},
			"compressed": {
				Type: schema.TypeBool,
				Required: false,
			},
			"location": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}