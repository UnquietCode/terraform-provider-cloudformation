// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchTemplateNetworkInterface(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplatePrivateIpAdd(),
				Required: false,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"device_index": {
				Type: schema.TypeInt,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateIpv6Add(),
				Required: false,
			},
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"interface_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"delete_on_termination": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}