// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayBasePathMappingType string = "AWS::ApiGateway::BasePathMapping"

func ResourceApiGatewayBasePathMapping() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayBasePathMappingRead,
		Create: resourceApiGatewayBasePathMappingCreate,
		Update: resourceApiGatewayBasePathMappingUpdate,
		Delete: resourceApiGatewayBasePathMappingDelete,
		CustomizeDiff: resourceApiGatewayBasePathMappingCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"base_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stage": {
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

func resourceApiGatewayBasePathMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayBasePathMappingType, ResourceApiGatewayBasePathMapping(), data, meta)
}

func resourceApiGatewayBasePathMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayBasePathMappingType, ResourceApiGatewayBasePathMapping(), data, meta)
}

func resourceApiGatewayBasePathMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayBasePathMappingType, ResourceApiGatewayBasePathMapping(), data, meta)
}

func resourceApiGatewayBasePathMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayBasePathMappingType, data, meta)
}

func resourceApiGatewayBasePathMappingCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayBasePathMappingType, data, meta)
}
