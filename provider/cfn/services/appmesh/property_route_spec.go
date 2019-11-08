// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRouteSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_route": {
				Type: schema.TypeList,
				Elem: propertyHttpRoute(),
				Required: false,
				MaxItems: 1,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: false,
			},
			"http2_route": {
				Type: schema.TypeList,
				Elem: propertyHttpRoute(),
				Required: false,
				MaxItems: 1,
			},
			"grpc_route": {
				Type: schema.TypeList,
				Elem: propertyGrpcRoute(),
				Required: false,
				MaxItems: 1,
			},
			"tcp_route": {
				Type: schema.TypeList,
				Elem: propertyTcpRoute(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}