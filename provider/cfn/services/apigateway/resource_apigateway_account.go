// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayAccountType string = "AWS::ApiGateway::Account"

var apiGatewayAccountProperties map[string]string = map[string]string{
	"cloud_watch_role_arn": "CloudWatchRoleArn",
}

func ResourceApiGatewayAccount() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayAccountExists,
		Read: resourceApiGatewayAccountRead,
		Create: resourceApiGatewayAccountCreate,
		Update: resourceApiGatewayAccountUpdate,
		Delete: resourceApiGatewayAccountDelete,
		CustomizeDiff: resourceApiGatewayAccountCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cloud_watch_role_arn": {
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

func resourceApiGatewayAccountExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayAccountRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayAccountType, ResourceApiGatewayAccount(), data, meta)
}

func resourceApiGatewayAccountCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayAccountType, ResourceApiGatewayAccount(), data, apiGatewayAccountProperties, meta)
}

func resourceApiGatewayAccountUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayAccountType, ResourceApiGatewayAccount(), data, apiGatewayAccountProperties, meta)
}

func resourceApiGatewayAccountDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayAccountType, data, meta)
}

func resourceApiGatewayAccountCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayAccountType, data, meta)
}
