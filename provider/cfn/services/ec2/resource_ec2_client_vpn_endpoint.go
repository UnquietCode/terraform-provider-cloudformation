// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2ClientVpnEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2ClientVpnEndpointExists,
		Read:   resourceEC2ClientVpnEndpointRead,
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
				Type: schema.TypeList,
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
	return plugin.ResourceRead("AWS::EC2::ClientVpnEndpoint", ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnEndpoint", ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnEndpoint", ResourceEC2ClientVpnEndpoint(), data, meta)
}

func resourceEC2ClientVpnEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnEndpoint", data, meta)
}

func resourceEC2ClientVpnEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}