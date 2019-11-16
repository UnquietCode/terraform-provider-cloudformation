// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2EgressOnlyInternetGatewayType string = "AWS::EC2::EgressOnlyInternetGateway"

func ResourceEC2EgressOnlyInternetGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2EgressOnlyInternetGatewayExists,
		Read: resourceEC2EgressOnlyInternetGatewayRead,
		Create: resourceEC2EgressOnlyInternetGatewayCreate,
		Update: resourceEC2EgressOnlyInternetGatewayUpdate,
		Delete: resourceEC2EgressOnlyInternetGatewayDelete,
		CustomizeDiff: resourceEC2EgressOnlyInternetGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"vpc_id": {
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

func resourceEC2EgressOnlyInternetGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2EgressOnlyInternetGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2EgressOnlyInternetGatewayType, ResourceEC2EgressOnlyInternetGateway(), data, meta)
}

func resourceEC2EgressOnlyInternetGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2EgressOnlyInternetGatewayType, ResourceEC2EgressOnlyInternetGateway(), data, meta)
}

func resourceEC2EgressOnlyInternetGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2EgressOnlyInternetGatewayType, ResourceEC2EgressOnlyInternetGateway(), data, meta)
}

func resourceEC2EgressOnlyInternetGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2EgressOnlyInternetGatewayType, data, meta)
}

func resourceEC2EgressOnlyInternetGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2EgressOnlyInternetGatewayType, data, meta)
}
