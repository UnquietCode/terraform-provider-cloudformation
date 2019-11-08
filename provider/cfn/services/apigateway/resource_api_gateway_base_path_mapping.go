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

func ResourceApiGatewayBasePathMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayBasePathMappingCreate,
		Read:   resourceApiGatewayBasePathMappingRead,
		Update: resourceApiGatewayBasePathMappingUpdate,
		Delete: resourceApiGatewayBasePathMappingDelete,

		Schema: map[string]*schema.Schema{
			"base_path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"stage": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayBasePathMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::BasePathMapping", data, meta)
}

func resourceApiGatewayBasePathMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::BasePathMapping", data, meta)
}

func resourceApiGatewayBasePathMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::BasePathMapping", data, meta)
}

func resourceApiGatewayBasePathMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::BasePathMapping", data, meta)
}