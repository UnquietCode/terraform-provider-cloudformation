// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIntegration() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegrationCreate,
		Read:   resourceIntegrationRead,
		Update: resourceIntegrationUpdate,
		Delete: resourceIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"template_selection_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"connection_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"integration_method": {
				Type: schema.TypeString,
				Required: false,
			},
			"passthrough_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"integration_uri": {
				Type: schema.TypeString,
				Required: false,
			},
			"credentials_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_templates": {
				Type: schema.TypeMap,
				Required: false,
			},
			"timeout_in_millis": {
				Type: schema.TypeInt,
				Required: false,
			},
			"content_handling_strategy": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"integration_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIntegrationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Integration", data, meta)
}

func resourceIntegrationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Integration", data, meta)
}

func resourceIntegrationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Integration", data, meta)
}

func resourceIntegrationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Integration", data, meta)
}