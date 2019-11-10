// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCPeeringConnectionCreate,
		Read:   resourceVPCPeeringConnectionRead,
		Update: resourceVPCPeeringConnectionUpdate,
		Delete: resourceVPCPeeringConnectionDelete,

		Schema: map[string]*schema.Schema{
			"peer_owner_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"peer_region": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"peer_role_arn": {
				Type: schema.TypeString,
				Required: false,
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
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCPeeringConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCPeeringConnection", data, meta)
}

func resourceVPCPeeringConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCPeeringConnection", data, meta)
}

func resourceVPCPeeringConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCPeeringConnection", data, meta)
}

func resourceVPCPeeringConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCPeeringConnection", data, meta)
}