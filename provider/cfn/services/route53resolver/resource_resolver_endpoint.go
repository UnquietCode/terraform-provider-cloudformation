// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResolverEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceResolverEndpointCreate,
		Read:   resourceResolverEndpointRead,
		Update: resourceResolverEndpointUpdate,
		Delete: resourceResolverEndpointDelete,

		Schema: map[string]*schema.Schema{
			"ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyResolverEndpointIpAddressRequest(),
				Required: true,
			},
			"direction": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceResolverEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceResolverEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceResolverEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceResolverEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}