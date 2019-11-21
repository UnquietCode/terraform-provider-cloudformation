// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaLayerVersionType string = "AWS::Lambda::LayerVersion"

func ResourceLambdaLayerVersion() *schema.Resource {
	return &schema.Resource{
		Read: resourceLambdaLayerVersionRead,
		Create: resourceLambdaLayerVersionCreate,
		Update: resourceLambdaLayerVersionUpdate,
		Delete: resourceLambdaLayerVersionDelete,
		CustomizeDiff: resourceLambdaLayerVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"compatible_runtimes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"license_info": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"layer_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"content": {
				Type: schema.TypeSet,
				Elem: propertyLayerVersionContent(),
				Required: true,
				MaxItems: 1,
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

func resourceLambdaLayerVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaLayerVersionType, ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lambdaLayerVersionType, ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lambdaLayerVersionType, ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lambdaLayerVersionType, data, meta)
}

func resourceLambdaLayerVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lambdaLayerVersionType, data, meta)
}
