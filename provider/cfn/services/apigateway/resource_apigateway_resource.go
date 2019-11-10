// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayResourceCreate,
		Read:   resourceApiGatewayResourceRead,
		Delete: resourceApiGatewayResourceDelete,

		Schema: map[string]*schema.Schema{
			"parent_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path_part": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Resource", data, meta)
}

func resourceApiGatewayResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Resource", data, meta)
}

func resourceApiGatewayResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Resource", data, meta)
}

func resourceApiGatewayResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Resource", data, meta)
}