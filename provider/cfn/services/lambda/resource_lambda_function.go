// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const lambdaFunctionType string = "AWS::Lambda::Function"

func ResourceLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Read: resourceLambdaFunctionRead,
		Create: resourceLambdaFunctionCreate,
		Update: resourceLambdaFunctionUpdate,
		Delete: resourceLambdaFunctionDelete,
		CustomizeDiff: resourceLambdaFunctionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"code": {
				Type: schema.TypeSet,
				Elem: propertyFunctionCode(),
				Required: true,
				MaxItems: 1,
			},
			"dead_letter_config": {
				Type: schema.TypeSet,
				Elem: propertyFunctionDeadLetterConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment": {
				Type: schema.TypeSet,
				Elem: propertyFunctionEnvironment(),
				Optional: true,
				MaxItems: 1,
			},
			"function_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"handler": {
				Type: schema.TypeString,
				Required: true,
			},
			"kms_key_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"layers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"memory_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"reserved_concurrent_executions": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"runtime": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tracing_config": {
				Type: schema.TypeSet,
				Elem: propertyFunctionTracingConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"vpc_config": {
				Type: schema.TypeSet,
				Elem: propertyFunctionVpcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLambdaFunctionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaFunctionType, ResourceLambdaFunction(), data, meta)
}

func resourceLambdaFunctionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lambdaFunctionType, ResourceLambdaFunction(), data, meta)
}

func resourceLambdaFunctionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lambdaFunctionType, ResourceLambdaFunction(), data, meta)
}

func resourceLambdaFunctionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lambdaFunctionType, data, meta)
}

func resourceLambdaFunctionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lambdaFunctionType, data, meta)
}
