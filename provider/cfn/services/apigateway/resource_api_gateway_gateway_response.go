// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayGatewayResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayGatewayResponseCreate,
		Read:   resourceApiGatewayGatewayResponseRead,
		Update: resourceApiGatewayGatewayResponseUpdate,
		Delete: resourceApiGatewayGatewayResponseDelete,

		Schema: map[string]*schema.Schema{
			"response_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"response_templates": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"response_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status_code": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayGatewayResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::GatewayResponse", data, meta)
}

func resourceApiGatewayGatewayResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::GatewayResponse", data, meta)
}

func resourceApiGatewayGatewayResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::GatewayResponse", data, meta)
}

func resourceApiGatewayGatewayResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::GatewayResponse", data, meta)
}