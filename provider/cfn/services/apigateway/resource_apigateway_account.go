// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayAccountType string = "AWS::ApiGateway::Account"

func ResourceApiGatewayAccount() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceApiGatewayAccountRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayAccountType, ResourceApiGatewayAccount(), data, meta)
}

func resourceApiGatewayAccountCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayAccountType, ResourceApiGatewayAccount(), data, meta)
}

func resourceApiGatewayAccountUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayAccountType, ResourceApiGatewayAccount(), data, meta)
}

func resourceApiGatewayAccountDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayAccountType, data, meta)
}

func resourceApiGatewayAccountCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayAccountType, data, meta)
}
