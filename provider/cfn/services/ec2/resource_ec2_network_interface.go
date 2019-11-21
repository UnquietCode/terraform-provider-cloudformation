// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkInterfaceType string = "AWS::EC2::NetworkInterface"

func ResourceEC2NetworkInterface() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2NetworkInterfaceRead,
		Create: resourceEC2NetworkInterfaceCreate,
		Update: resourceEC2NetworkInterfaceUpdate,
		Delete: resourceEC2NetworkInterfaceDelete,
		CustomizeDiff: resourceEC2NetworkInterfaceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeSet,
				Elem: propertyNetworkInterfaceInstanceIpv6Address(),
				Optional: true,
				MaxItems: 1,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2NetworkInterfaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkInterfaceType, ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NetworkInterfaceType, ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NetworkInterfaceType, ResourceEC2NetworkInterface(), data, meta)
}

func resourceEC2NetworkInterfaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NetworkInterfaceType, data, meta)
}

func resourceEC2NetworkInterfaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NetworkInterfaceType, data, meta)
}
