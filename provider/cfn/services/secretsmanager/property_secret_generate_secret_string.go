// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html

package secretsmanager

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySecretGenerateSecretString(extras...string) *schema.Resource {
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
			"exclude_uppercase": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"require_each_included_type": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_space": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"exclude_characters": {
				Type: schema.TypeString,
				Optional: true,
			},
			"generate_string_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"password_length": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"exclude_punctuation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"exclude_lowercase": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"secret_string_template": {
				Type: schema.TypeString,
				Optional: true,
			},
			"exclude_numbers": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
