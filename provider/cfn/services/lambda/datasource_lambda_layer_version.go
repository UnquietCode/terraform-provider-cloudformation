// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceLambdaLayerVersion() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaLayerVersionRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceLambdaLayerVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaLayerVersionType, DatasourceLambdaLayerVersion(), data, meta)
}
