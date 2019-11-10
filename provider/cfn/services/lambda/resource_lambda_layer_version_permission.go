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

func ResourceLambdaLayerVersionPermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaLayerVersionPermissionCreate,
		Read:   resourceLambdaLayerVersionPermissionRead,
		Update: resourceLambdaLayerVersionPermissionUpdate,
		Delete: resourceLambdaLayerVersionPermissionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"layer_version_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLambdaLayerVersionPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::LayerVersionPermission", data, meta)
}

func resourceLambdaLayerVersionPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::LayerVersionPermission", data, meta)
}

func resourceLambdaLayerVersionPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::LayerVersionPermission", data, meta)
}

func resourceLambdaLayerVersionPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::LayerVersionPermission", data, meta)
}