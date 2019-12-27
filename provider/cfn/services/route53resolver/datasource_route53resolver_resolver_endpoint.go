// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html

package route53resolver

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53ResolverResolverEndpointType string = "AWS::Route53Resolver::ResolverEndpoint"

func DatasourceRoute53ResolverResolverEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoute53ResolverResolverEndpointRead,
		
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
			"tags": misc.PropertyTags(),
			"name": {
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

func datasourceRoute53ResolverResolverEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53ResolverResolverEndpointType, DatasourceRoute53ResolverResolverEndpoint(), data, meta)
}
