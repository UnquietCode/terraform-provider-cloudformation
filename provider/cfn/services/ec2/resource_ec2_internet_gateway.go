// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2InternetGatewayType string = "AWS::EC2::InternetGateway"

var eC2InternetGatewayProperties map[string]string = map[string]string{
	"tags": "Tags",
}

func ResourceEC2InternetGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2InternetGatewayExists,
		Read: resourceEC2InternetGatewayRead,
		Create: resourceEC2InternetGatewayCreate,
		Update: resourceEC2InternetGatewayUpdate,
		Delete: resourceEC2InternetGatewayDelete,
		CustomizeDiff: resourceEC2InternetGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceEC2InternetGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2InternetGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, meta)
}

func resourceEC2InternetGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, eC2InternetGatewayProperties, meta)
}

func resourceEC2InternetGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2InternetGatewayType, ResourceEC2InternetGateway(), data, eC2InternetGatewayProperties, meta)
}

func resourceEC2InternetGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2InternetGatewayType, data, meta)
}

func resourceEC2InternetGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2InternetGatewayType, data, meta)
}
