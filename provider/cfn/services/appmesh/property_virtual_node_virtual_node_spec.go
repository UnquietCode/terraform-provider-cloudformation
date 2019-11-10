// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyVirtualNodeVirtualNodeSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeLogging(),
				Required: false,
				MaxItems: 1,
			},
			"backends": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeBackend(),
				Required: false,
			},
			"listeners": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeListener(),
				Required: false,
			},
			"service_discovery": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeServiceDiscovery(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}