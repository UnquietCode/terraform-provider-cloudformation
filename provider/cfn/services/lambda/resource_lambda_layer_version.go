// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaLayerVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLambdaLayerVersionExists,
		Read:   resourceLambdaLayerVersionRead,
		Create: resourceLambdaLayerVersionCreate,
		Update: resourceLambdaLayerVersionUpdate,
		Delete: resourceLambdaLayerVersionDelete,
		
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
				Type: schema.TypeList,
				Elem: propertyLayerVersionContent(),
				Required: true,
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

func resourceLambdaLayerVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLambdaLayerVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::LayerVersion", ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::LayerVersion", ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::LayerVersion", ResourceLambdaLayerVersion(), data, meta)
}

func resourceLambdaLayerVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::LayerVersion", data, meta)
}