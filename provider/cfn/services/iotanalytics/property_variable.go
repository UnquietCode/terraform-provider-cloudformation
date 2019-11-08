// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyVariable() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dataset_content_version_value": {
				Type: schema.TypeList,
				Elem: propertyDatasetContentVersionValue(),
				Required: false,
				MaxItems: 1,
			},
			"double_value": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"output_file_uri_value": {
				Type: schema.TypeList,
				Elem: propertyOutputFileUriValue(),
				Required: false,
				MaxItems: 1,
			},
			"variable_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"string_value": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}