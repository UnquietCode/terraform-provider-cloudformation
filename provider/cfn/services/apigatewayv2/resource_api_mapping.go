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

func ResourceApiMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiMappingCreate,
		Read:   resourceApiMappingRead,
		Update: resourceApiMappingUpdate,
		Delete: resourceApiMappingDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stage": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_mapping_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::ApiMapping", data, meta)
}

func resourceApiMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::ApiMapping", data, meta)
}

func resourceApiMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::ApiMapping", data, meta)
}

func resourceApiMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::ApiMapping", data, meta)
}