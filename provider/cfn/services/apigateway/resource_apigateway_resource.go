// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceApiGatewayResourceExists,
		Read: resourceApiGatewayResourceRead,
		Create: resourceApiGatewayResourceCreate,
		Update: resourceApiGatewayResourceUpdate,
		Delete: resourceApiGatewayResourceDelete,
		CustomizeDiff: resourceApiGatewayResourceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"parent_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"path_part": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
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

func resourceApiGatewayResourceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Resource", ResourceApiGatewayResource(), data, meta)
}

func resourceApiGatewayResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Resource", ResourceApiGatewayResource(), data, meta)
}

func resourceApiGatewayResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Resource", ResourceApiGatewayResource(), data, meta)
}

func resourceApiGatewayResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Resource", data, meta)
}

func resourceApiGatewayResourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ApiGateway::Resource", data, meta)
}

