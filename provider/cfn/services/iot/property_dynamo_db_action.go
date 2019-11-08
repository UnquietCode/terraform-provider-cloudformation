// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDynamoDBAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hash_key_field": {
				Type: schema.TypeString,
				Required: true,
			},
			"hash_key_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"hash_key_value": {
				Type: schema.TypeString,
				Required: true,
			},
			"payload_field": {
				Type: schema.TypeString,
				Required: false,
			},
			"range_key_field": {
				Type: schema.TypeString,
				Required: false,
			},
			"range_key_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"range_key_value": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}