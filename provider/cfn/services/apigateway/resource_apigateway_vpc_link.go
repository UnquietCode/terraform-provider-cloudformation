// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayVpcLinkType string = "AWS::ApiGateway::VpcLink"

func ResourceApiGatewayVpcLink() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayVpcLinkRead,
		Create: resourceApiGatewayVpcLinkCreate,
		Update: resourceApiGatewayVpcLinkUpdate,
		Delete: resourceApiGatewayVpcLinkDelete,
		CustomizeDiff: resourceApiGatewayVpcLinkCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func resourceApiGatewayVpcLinkRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayVpcLinkType, ResourceApiGatewayVpcLink(), data, meta)
}

func resourceApiGatewayVpcLinkCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayVpcLinkType, ResourceApiGatewayVpcLink(), data, meta)
}

func resourceApiGatewayVpcLinkUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayVpcLinkType, ResourceApiGatewayVpcLink(), data, meta)
}

func resourceApiGatewayVpcLinkDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayVpcLinkType, data, meta)
}

func resourceApiGatewayVpcLinkCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayVpcLinkType, data, meta)
}
