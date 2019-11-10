// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaPermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaPermissionCreate,
		Read:   resourceLambdaPermissionRead,
		Update: resourceLambdaPermissionUpdate,
		Delete: resourceLambdaPermissionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"event_source_token": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_account": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLambdaPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Permission", data, meta)
}

func resourceLambdaPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Permission", data, meta)
}

func resourceLambdaPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Permission", data, meta)
}

func resourceLambdaPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Permission", data, meta)
}