// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaFunctionCreate,
		Read:   resourceLambdaFunctionRead,
		Update: resourceLambdaFunctionUpdate,
		Delete: resourceLambdaFunctionDelete,

		Schema: map[string]*schema.Schema{
			"code": {
				Type: schema.TypeList,
				Elem: propertyFunctionCode(),
				Required: true,
				MaxItems: 1,
			},
			"dead_letter_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionDeadLetterConfig(),
				Required: false,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyFunctionEnvironment(),
				Required: false,
				MaxItems: 1,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"handler": {
				Type: schema.TypeString,
				Required: true,
			},
			"kms_key_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"layers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"memory_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"reserved_concurrent_executions": {
				Type: schema.TypeInt,
				Required: false,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"runtime": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"tracing_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionTracingConfig(),
				Required: false,
				MaxItems: 1,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionVpcConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceLambdaFunctionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Function", data, meta)
}

func resourceLambdaFunctionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Function", data, meta)
}

func resourceLambdaFunctionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Function", data, meta)
}

func resourceLambdaFunctionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Function", data, meta)
}