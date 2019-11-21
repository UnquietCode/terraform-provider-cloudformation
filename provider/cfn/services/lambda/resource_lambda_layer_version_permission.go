// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaLayerVersionPermissionType string = "AWS::Lambda::LayerVersionPermission"

func ResourceLambdaLayerVersionPermission() *schema.Resource {
	return &schema.Resource{
		Read: resourceLambdaLayerVersionPermissionRead,
		Create: resourceLambdaLayerVersionPermissionCreate,
		Update: resourceLambdaLayerVersionPermissionUpdate,
		Delete: resourceLambdaLayerVersionPermissionDelete,
		CustomizeDiff: resourceLambdaLayerVersionPermissionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
			},
			"layer_version_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"organization_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
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

func resourceLambdaLayerVersionPermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaLayerVersionPermissionType, ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lambdaLayerVersionPermissionType, ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lambdaLayerVersionPermissionType, ResourceLambdaLayerVersionPermission(), data, meta)
}

func resourceLambdaLayerVersionPermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lambdaLayerVersionPermissionType, data, meta)
}

func resourceLambdaLayerVersionPermissionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lambdaLayerVersionPermissionType, data, meta)
}
