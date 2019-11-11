// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2NetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2NetworkInterfaceCreate,
		Read:   resourceEC2NetworkInterfaceRead,
		Update: resourceEC2NetworkInterfaceUpdate,
		Delete: resourceEC2NetworkInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"primary_private_ip_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"secondary_private_ip_addresses": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_set": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"interface_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyNetworkInterfaceInstanceIpv6Address(),
				Optional: true,
				MaxItems: 1,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"private_ip_addresses": {
				Type: schema.TypeSet,
				Elem: propertyNetworkInterfacePrivateIpAddressSpecification(),
				Optional: true,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"source_dest_check": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func resourceEC2NetworkInterfaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterface", ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterface", ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterface", ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterface", data, meta)
}