// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplatePrivateIpAdd(),
				Optional: true,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"device_index": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateIpv6Add(),
				Optional: true,
			},
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"interface_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"delete_on_termination": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
