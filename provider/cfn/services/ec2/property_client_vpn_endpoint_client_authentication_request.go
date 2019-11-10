// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClientVpnEndpointClientAuthenticationRequest() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mutual_authentication": {
				Type: schema.TypeList,
				Elem: propertyClientVpnEndpointCertificateAuthenticationRequest(),
				Required: false,
				MaxItems: 1,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"active_directory": {
				Type: schema.TypeList,
				Elem: propertyClientVpnEndpointDirectoryServiceAuthenticationRequest(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}