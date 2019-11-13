// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaPermission() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLambdaPermissionExists,
		Read:   resourceLambdaPermissionRead,
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
			},
		},
	}
}

func resourceLambdaPermissionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLambdaPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Permission", ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Permission", ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Permission", ResourceLambdaPermission(), data, meta)
}

func resourceLambdaPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Permission", data, meta)
}

func resourceLambdaPermissionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}