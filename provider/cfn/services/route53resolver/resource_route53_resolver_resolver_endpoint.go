// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53ResolverResolverEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoute53ResolverResolverEndpointCreate,
		Read:   resourceRoute53ResolverResolverEndpointRead,
		Update: resourceRoute53ResolverResolverEndpointUpdate,
		Delete: resourceRoute53ResolverResolverEndpointDelete,

		Schema: map[string]*schema.Schema{
			"ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyIpAddressRequest(),
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

func resourceRoute53ResolverResolverEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceRoute53ResolverResolverEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceRoute53ResolverResolverEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}

func resourceRoute53ResolverResolverEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverEndpoint", data, meta)
}