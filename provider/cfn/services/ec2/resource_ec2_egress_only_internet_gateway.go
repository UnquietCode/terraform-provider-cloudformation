// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2EgressOnlyInternetGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2EgressOnlyInternetGatewayCreate,
		Read:   resourceEC2EgressOnlyInternetGatewayRead,
		Update: resourceEC2EgressOnlyInternetGatewayUpdate,
		Delete: resourceEC2EgressOnlyInternetGatewayDelete,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2EgressOnlyInternetGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EgressOnlyInternetGateway", data, meta)
}

func resourceEC2EgressOnlyInternetGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EgressOnlyInternetGateway", data, meta)
}

func resourceEC2EgressOnlyInternetGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EgressOnlyInternetGateway", data, meta)
}

func resourceEC2EgressOnlyInternetGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EgressOnlyInternetGateway", data, meta)
}