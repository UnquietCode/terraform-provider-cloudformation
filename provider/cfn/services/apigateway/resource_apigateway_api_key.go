// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayApiKeyType string = "AWS::ApiGateway::ApiKey"

var apiGatewayApiKeyProperties map[string]string = map[string]string{
	"customer_id": "CustomerId",
	"description": "Description",
	"enabled": "Enabled",
	"generate_distinct_id": "GenerateDistinctId",
	"name": "Name",
	"stage_keys": "StageKeys",
	"value": "Value",
}

func ResourceApiGatewayApiKey() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayApiKeyExists,
		Read: resourceApiGatewayApiKeyRead,
		Create: resourceApiGatewayApiKeyCreate,
		Update: resourceApiGatewayApiKeyUpdate,
		Delete: resourceApiGatewayApiKeyDelete,
		CustomizeDiff: resourceApiGatewayApiKeyCustomizeDiff,
		
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
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stage_keys": {
				Type: schema.TypeSet,
				Elem: propertyApiKeyStageKey(),
				Optional: true,
			},
			"value": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayApiKeyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayApiKeyType, ResourceApiGatewayApiKey(), data, meta)
}

func resourceApiGatewayApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayApiKeyType, ResourceApiGatewayApiKey(), data, apiGatewayApiKeyProperties, meta)
}

func resourceApiGatewayApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayApiKeyType, ResourceApiGatewayApiKey(), data, apiGatewayApiKeyProperties, meta)
}

func resourceApiGatewayApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayApiKeyType, data, meta)
}

func resourceApiGatewayApiKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayApiKeyType, data, meta)
}
