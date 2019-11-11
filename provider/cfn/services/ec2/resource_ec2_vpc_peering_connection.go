// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPCPeeringConnectionCreate,
		Read:   resourceEC2VPCPeeringConnectionRead,
		Update: resourceEC2VPCPeeringConnectionUpdate,
		Delete: resourceEC2VPCPeeringConnectionDelete,

		Schema: map[string]*schema.Schema{
			"peer_owner_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_region": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_role_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
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

func resourceEC2VPCPeeringConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCPeeringConnection", ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCPeeringConnection", ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCPeeringConnection", ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCPeeringConnection", data, meta)
}