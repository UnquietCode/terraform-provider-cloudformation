// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaFunctionType string = "AWS::Lambda::Function"

func DatasourceLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaFunctionRead,
		
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
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
				Elem: propertyFunctionTracingConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionVpcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceLambdaFunctionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaFunctionType, DatasourceLambdaFunction(), data, meta)
}
