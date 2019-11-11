// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaLayerVersionPermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaLayerVersionPermissionCreate,
		Read:   resourceLambdaLayerVersionPermissionRead,
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
				Optional: true,
				ForceNew: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLambdaLayerVersionPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::LayerVersionPermission", ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::LayerVersionPermission", ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::LayerVersionPermission", ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::LayerVersionPermission", data, meta)
}