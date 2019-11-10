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

func ResourceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfaceCreate,
		Read:   resourceNetworkInterfaceRead,
		Update: resourceNetworkInterfaceUpdate,
		Delete: resourceNetworkInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"group_set": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"interface_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyNetworkInterfaceInstanceIpv6Address(),
				Required: false,
				MaxItems: 1,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"private_ip_addresses": {
				Type: schema.TypeSet,
				Elem: propertyNetworkInterfacePrivateIpAddressSpecification(),
				Required: false,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"source_dest_check": {
				Type: schema.TypeBool,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceNetworkInterfaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterface", data, meta)
}

func resourceNetworkInterfaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterface", data, meta)
}

func resourceNetworkInterfaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterface", data, meta)
}

func resourceNetworkInterfaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterface", data, meta)
}