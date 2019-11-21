// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayRequestValidatorType string = "AWS::ApiGateway::RequestValidator"

func ResourceApiGatewayRequestValidator() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayRequestValidatorRead,
		Create: resourceApiGatewayRequestValidatorCreate,
		Update: resourceApiGatewayRequestValidatorUpdate,
		Delete: resourceApiGatewayRequestValidatorDelete,
		CustomizeDiff: resourceApiGatewayRequestValidatorCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"validate_request_body": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"validate_request_parameters": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceApiGatewayRequestValidatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayRequestValidatorType, ResourceApiGatewayRequestValidator(), data, meta)
}

func resourceApiGatewayRequestValidatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayRequestValidatorType, ResourceApiGatewayRequestValidator(), data, meta)
}

func resourceApiGatewayRequestValidatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayRequestValidatorType, ResourceApiGatewayRequestValidator(), data, meta)
}

func resourceApiGatewayRequestValidatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayRequestValidatorType, data, meta)
}

func resourceApiGatewayRequestValidatorCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayRequestValidatorType, data, meta)
}
