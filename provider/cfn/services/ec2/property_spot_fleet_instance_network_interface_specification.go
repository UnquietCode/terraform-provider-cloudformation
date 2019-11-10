// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySpotFleetInstanceNetworkInterfaceSpecification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Required: false,
			},
			"delete_on_termination": {
				Type: schema.TypeBool,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"device_index": {
				Type: schema.TypeInt,
				Required: false,
			},
			"groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ipv6_addresses": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetInstanceIpv6Address(),
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_addresses": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetPrivateIpAddressSpecification(),
				Required: false,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}