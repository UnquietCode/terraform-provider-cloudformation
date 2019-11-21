// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaPermissionType string = "AWS::Lambda::Permission"

func ResourceLambdaPermission() *schema.Resource {
	return &schema.Resource{
		Read: resourceLambdaPermissionRead,
		Create: resourceLambdaPermissionCreate,
		Update: resourceLambdaPermissionUpdate,
		Delete: resourceLambdaPermissionDelete,
		CustomizeDiff: resourceLambdaPermissionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
			},
			"event_source_token": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_account": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceLambdaPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaPermissionType, ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lambdaPermissionType, ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lambdaPermissionType, ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lambdaPermissionType, data, meta)
}

func resourceLambdaPermissionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lambdaPermissionType, data, meta)
}
