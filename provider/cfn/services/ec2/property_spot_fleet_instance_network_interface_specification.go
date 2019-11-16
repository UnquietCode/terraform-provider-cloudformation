// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var spotFleetInstanceNetworkInterfaceSpecificationProperties map[string]string = map[string]string{
	"associate_public_ip_address": "AssociatePublicIpAddress",
	"delete_on_termination": "DeleteOnTermination",
	"description": "Description",
	"device_index": "DeviceIndex",
	"groups": "Groups",
	"ipv6_address_count": "Ipv6AddressCount",
	"ipv6_addresses": "Ipv6Addresses",
	"network_interface_id": "NetworkInterfaceId",
	"private_ip_addresses": "PrivateIpAddresses",
	"secondary_private_ip_address_count": "SecondaryPrivateIpAddressCount",
	"subnet_id": "SubnetId",
}

func propertySpotFleetInstanceNetworkInterfaceSpecification(extras...string) *schema.Resource {
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
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"delete_on_termination": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"device_index": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetInstanceIpv6Address(),
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_ip_addresses": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetPrivateIpAddressSpecification(),
				Optional: true,
			},
			"secondary_private_ip_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
