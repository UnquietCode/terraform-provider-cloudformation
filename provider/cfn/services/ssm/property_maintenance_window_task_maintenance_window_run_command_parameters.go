// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html

package ssm

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMaintenanceWindowTaskMaintenanceWindowRunCommandParameters(extras...string) *schema.Resource {
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
			"timeout_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"comment": {
				Type: schema.TypeString,
				Optional: true,
			},
			"output_s3_key_prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"document_hash_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_config": {
				Type: schema.TypeSet,
				Elem: propertyMaintenanceWindowTaskNotificationConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"output_s3_bucket_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"document_hash": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
