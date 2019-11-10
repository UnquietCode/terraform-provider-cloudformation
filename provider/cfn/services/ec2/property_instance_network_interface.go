// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceNetworkInterface() *schema.Resource {
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
				Type: schema.TypeString,
				Required: true,
			},
			"group_set": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyInstanceInstanceIpv6Address(),
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyInstancePrivateIpAddressSpecification(),
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