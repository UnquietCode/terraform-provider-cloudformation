// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2ApiMappingType string = "AWS::ApiGatewayV2::ApiMapping"

func ResourceApiGatewayV2ApiMapping() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayV2ApiMappingExists,
		Read: resourceApiGatewayV2ApiMappingRead,
		Create: resourceApiGatewayV2ApiMappingCreate,
		Update: resourceApiGatewayV2ApiMappingUpdate,
		Delete: resourceApiGatewayV2ApiMappingDelete,
		CustomizeDiff: resourceApiGatewayV2ApiMappingCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_mapping_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayV2ApiMappingExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2ApiMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2ApiMappingType, ResourceApiGatewayV2ApiMapping(), data, meta)
}

func resourceApiGatewayV2ApiMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2ApiMappingType, ResourceApiGatewayV2ApiMapping(), data, meta)
}

func resourceApiGatewayV2ApiMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2ApiMappingType, ResourceApiGatewayV2ApiMapping(), data, meta)
}

func resourceApiGatewayV2ApiMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2ApiMappingType, data, meta)
}

func resourceApiGatewayV2ApiMappingCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2ApiMappingType, data, meta)
}
