// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnEndpointType string = "AWS::EC2::ClientVpnEndpoint"

func ResourceEC2ClientVpnEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2ClientVpnEndpointExists,
		Read: resourceEC2ClientVpnEndpointRead,
		Create: resourceEC2ClientVpnEndpointCreate,
		Update: resourceEC2ClientVpnEndpointUpdate,
		Delete: resourceEC2ClientVpnEndpointDelete,
		CustomizeDiff: resourceEC2ClientVpnEndpointCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"connection_log_options": {
				Type: schema.TypeSet,
				Elem: propertyClientVpnEndpointConnectionLogOptions(),
				Required: true,
				MaxItems: 1,
			},
			"split_tunnel": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyClientVpnEndpointTagSpecification(),
				Optional: true,
			},
			"authentication_options": {
				Type: schema.TypeList,
				Elem: propertyClientVpnEndpointClientAuthenticationRequest(),
				Required: true,
			},
			"server_certificate_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"transport_protocol": {
				Type: schema.TypeString,
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

func resourceEC2ClientVpnEndpointExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2ClientVpnEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnEndpointType, ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2ClientVpnEndpointType, ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2ClientVpnEndpointType, ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2ClientVpnEndpointType, data, meta)
}

func resourceEC2ClientVpnEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2ClientVpnEndpointType, data, meta)
}
