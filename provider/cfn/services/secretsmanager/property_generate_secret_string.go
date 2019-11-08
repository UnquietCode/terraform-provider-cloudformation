// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGenerateSecretString() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exclude_uppercase": {
				Type: schema.TypeBool,
				Required: false,
			},
			"require_each_included_type": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_space": {
				Type: schema.TypeBool,
				Required: false,
			},
			"exclude_characters": {
				Type: schema.TypeString,
				Required: false,
			},
			"generate_string_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"password_length": {
				Type: schema.TypeInt,
				Required: false,
			},
			"exclude_punctuation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"exclude_lowercase": {
				Type: schema.TypeBool,
				Required: false,
			},
			"secret_string_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"exclude_numbers": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}