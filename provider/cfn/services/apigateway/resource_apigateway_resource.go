// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayResourceType string = "AWS::ApiGateway::Resource"

var apiGatewayResourceProperties map[string]string = map[string]string{
	"parent_id": "ParentId",
	"path_part": "PathPart",
	"rest_api_id": "RestApiId",
}

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
	return plugin.ResourceRead(apiGatewayResourceType, ResourceApiGatewayResource(), data, meta)
}

func resourceApiGatewayResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayResourceType, ResourceApiGatewayResource(), data, apiGatewayResourceProperties, meta)
}

func resourceApiGatewayResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayResourceType, ResourceApiGatewayResource(), data, apiGatewayResourceProperties, meta)
}

func resourceApiGatewayResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayResourceType, data, meta)
}

func resourceApiGatewayResourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayResourceType, data, meta)
}
