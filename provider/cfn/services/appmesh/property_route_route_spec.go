// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html

package appmesh

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var routeRouteSpecProperties map[string]string = map[string]string{
	"http_route": "HttpRoute",
	"priority": "Priority",
	"http2_route": "Http2Route",
	"grpc_route": "GrpcRoute",
	"tcp_route": "TcpRoute",
}

func propertyRouteRouteSpec(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_route": {
				Type: schema.TypeList,
				Elem: propertyRouteHttpRoute(),
				Optional: true,
				MaxItems: 1,
			},
			"priority": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"http2_route": {
				Type: schema.TypeList,
				Elem: propertyRouteHttpRoute(),
				Optional: true,
				MaxItems: 1,
			},
			"grpc_route": {
				Type: schema.TypeList,
				Elem: propertyRouteGrpcRoute(),
				Optional: true,
				MaxItems: 1,
			},
			"tcp_route": {
				Type: schema.TypeList,
				Elem: propertyRouteTcpRoute(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
