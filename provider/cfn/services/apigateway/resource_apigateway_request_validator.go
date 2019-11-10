// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayRequestValidator() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayRequestValidatorCreate,
		Read:   resourceApiGatewayRequestValidatorRead,
		Update: resourceApiGatewayRequestValidatorUpdate,
		Delete: resourceApiGatewayRequestValidatorDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"validate_request_body": {
				Type: schema.TypeBool,
				Required: false,
			},
			"validate_request_parameters": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayRequestValidatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::RequestValidator", data, meta)
}

func resourceApiGatewayRequestValidatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::RequestValidator", data, meta)
}

func resourceApiGatewayRequestValidatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::RequestValidator", data, meta)
}

func resourceApiGatewayRequestValidatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::RequestValidator", data, meta)
}