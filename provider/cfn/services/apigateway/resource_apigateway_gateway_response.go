// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayGatewayResponse() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayGatewayResponseExists,
		Read: resourceApiGatewayGatewayResponseRead,
		Create: resourceApiGatewayGatewayResponseCreate,
		Update: resourceApiGatewayGatewayResponseUpdate,
		Delete: resourceApiGatewayGatewayResponseDelete,
		CustomizeDiff: resourceApiGatewayGatewayResponseCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"response_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"response_templates": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"response_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"status_code": {
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

func resourceApiGatewayGatewayResponseExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayGatewayResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::GatewayResponse", ResourceApiGatewayGatewayResponse(), data, meta)
}

func resourceApiGatewayGatewayResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::GatewayResponse", ResourceApiGatewayGatewayResponse(), data, meta)
}

func resourceApiGatewayGatewayResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::GatewayResponse", ResourceApiGatewayGatewayResponse(), data, meta)
}

func resourceApiGatewayGatewayResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::GatewayResponse", data, meta)
}

func resourceApiGatewayGatewayResponseCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
