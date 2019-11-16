// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53ResolverResolverEndpointType string = "AWS::Route53Resolver::ResolverEndpoint"

var route53ResolverResolverEndpointProperties map[string]string = map[string]string{
	"ip_addresses": "IpAddresses",
	"direction": "Direction",
	"security_group_ids": "SecurityGroupIds",
	"tags": "Tags",
	"name": "Name",
}

func ResourceRoute53ResolverResolverEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoute53ResolverResolverEndpointExists,
		Read: resourceRoute53ResolverResolverEndpointRead,
		Create: resourceRoute53ResolverResolverEndpointCreate,
		Update: resourceRoute53ResolverResolverEndpointUpdate,
		Delete: resourceRoute53ResolverResolverEndpointDelete,
		CustomizeDiff: resourceRoute53ResolverResolverEndpointCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"ip_addresses": {
				Type: schema.TypeList,
				Elem: propertyResolverEndpointIpAddressRequest(),
				Required: true,
			},
			"direction": {
				Type: schema.TypeString,
				Required: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
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

func resourceRoute53ResolverResolverEndpointExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoute53ResolverResolverEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53ResolverResolverEndpointType, ResourceRoute53ResolverResolverEndpoint(), data, meta)
}

func resourceRoute53ResolverResolverEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53ResolverResolverEndpointType, ResourceRoute53ResolverResolverEndpoint(), data, route53ResolverResolverEndpointProperties, meta)
}

func resourceRoute53ResolverResolverEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53ResolverResolverEndpointType, ResourceRoute53ResolverResolverEndpoint(), data, route53ResolverResolverEndpointProperties, meta)
}

func resourceRoute53ResolverResolverEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53ResolverResolverEndpointType, data, meta)
}

func resourceRoute53ResolverResolverEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53ResolverResolverEndpointType, data, meta)
}
