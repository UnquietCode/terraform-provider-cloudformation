// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayApiKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayApiKeyCreate,
		Read:   resourceApiGatewayApiKeyRead,
		Update: resourceApiGatewayApiKeyUpdate,
		Delete: resourceApiGatewayApiKeyDelete,

		Schema: map[string]*schema.Schema{
			"customer_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"generate_distinct_id": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"stage_keys": {
				Type: schema.TypeSet,
				Elem: propertyApiKeyStageKey(),
				Optional: true,
			},
			"value": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::ApiKey", ResourceApiGatewayApiKey(), data, meta)
}

func resourceApiGatewayApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::ApiKey", ResourceApiGatewayApiKey(), data, meta)
}

func resourceApiGatewayApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::ApiKey", ResourceApiGatewayApiKey(), data, meta)
}

func resourceApiGatewayApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::ApiKey", data, meta)
}