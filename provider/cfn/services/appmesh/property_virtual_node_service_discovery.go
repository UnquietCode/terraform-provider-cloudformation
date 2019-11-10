// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyVirtualNodeServiceDiscovery() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeDnsServiceDiscovery(),
				Required: false,
				MaxItems: 1,
			},
			"aws_cloud_map": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeAwsCloudMapServiceDiscovery(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}