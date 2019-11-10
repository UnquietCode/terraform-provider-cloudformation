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

func ResourceApiGatewayVpcLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayVpcLinkCreate,
		Read:   resourceApiGatewayVpcLinkRead,
		Update: resourceApiGatewayVpcLinkUpdate,
		Delete: resourceApiGatewayVpcLinkDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"target_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApiGatewayVpcLinkCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::VpcLink", data, meta)
}

func resourceApiGatewayVpcLinkRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::VpcLink", data, meta)
}

func resourceApiGatewayVpcLinkUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::VpcLink", data, meta)
}

func resourceApiGatewayVpcLinkDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::VpcLink", data, meta)
}