// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnEndpointType string = "AWS::EC2::ClientVpnEndpoint"

func DatasourceEC2ClientVpnEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2ClientVpnEndpointRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2ClientVpnEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnEndpointType, DatasourceEC2ClientVpnEndpoint(), data, meta)
}
