// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTarget() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"ecs_parameters": {
				Type: schema.TypeList,
				Elem: propertyEcsParameters(),
				Required: false,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"input": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_transformer": {
				Type: schema.TypeList,
				Elem: propertyInputTransformer(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_parameters": {
				Type: schema.TypeList,
				Elem: propertyKinesisParameters(),
				Required: false,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"run_command_parameters": {
				Type: schema.TypeList,
				Elem: propertyRunCommandParameters(),
				Required: false,
				MaxItems: 1,
			},
			"sqs_parameters": {
				Type: schema.TypeList,
				Elem: propertySqsParameters(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}