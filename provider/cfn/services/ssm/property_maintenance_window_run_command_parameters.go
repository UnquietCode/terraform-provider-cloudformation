// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMaintenanceWindowRunCommandParameters() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"timeout_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"comment": {
				Type: schema.TypeString,
				Required: false,
			},
			"output_s3_key_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"document_hash_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_config": {
				Type: schema.TypeList,
				Elem: propertyNotificationConfig(),
				Required: false,
				MaxItems: 1,
			},
			"output_s3_bucket_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"document_hash": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}