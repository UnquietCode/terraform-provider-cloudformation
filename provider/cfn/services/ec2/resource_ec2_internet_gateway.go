// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2InternetGatewayType string = "AWS::EC2::InternetGateway"

func ResourceEC2InternetGateway() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2InternetGatewayRead,
		Create: resourceEC2InternetGatewayCreate,
		Update: resourceEC2InternetGatewayUpdate,
		Delete: resourceEC2InternetGatewayDelete,
		CustomizeDiff: resourceEC2InternetGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2InternetGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, meta)
}

func resourceEC2InternetGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, meta)
}

func resourceEC2InternetGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, meta)
}

func resourceEC2InternetGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2InternetGatewayType, data, meta)
}

func resourceEC2InternetGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2InternetGatewayType, data, meta)
}
