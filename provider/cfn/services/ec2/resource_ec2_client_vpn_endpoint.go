// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2ClientVpnEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2ClientVpnEndpointCreate,
		Read:   resourceEC2ClientVpnEndpointRead,
		Update: resourceEC2ClientVpnEndpointUpdate,
		Delete: resourceEC2ClientVpnEndpointDelete,

		Schema: map[string]*schema.Schema{
			"client_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_log_options": {
				Type: schema.TypeList,
				Elem: propertyConnectionLogOptions(),
				Required: true,
				MaxItems: 1,
			},
			"split_tunnel": {
				Type: schema.TypeBool,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyTagSpecification(),
				Required: false,
				ForceNew: true,
			},
			"authentication_options": {
				Type: schema.TypeList,
				Elem: propertyClientAuthenticationRequest(),
				Required: true,
				ForceNew: true,
			},
			"server_certificate_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"transport_protocol": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2ClientVpnEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnEndpoint", data, meta)
}

func resourceEC2ClientVpnEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnEndpoint", data, meta)
}

func resourceEC2ClientVpnEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnEndpoint", data, meta)
}

func resourceEC2ClientVpnEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnEndpoint", data, meta)
}