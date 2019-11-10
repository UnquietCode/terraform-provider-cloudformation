// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		return nil
	}
	
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
				Elem: propertyMaintenanceWindowTaskNotificationConfig(),
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