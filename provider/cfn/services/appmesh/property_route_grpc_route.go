// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRouteGrpcRoute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList,
				Elem: propertyRouteGrpcRouteAction(),
				Required: true,
				MaxItems: 1,
			},
			"retry_policy": {
				Type: schema.TypeList,
				Elem: propertyRouteGrpcRetryPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"match": {
				Type: schema.TypeList,
				Elem: propertyRouteGrpcRouteMatch(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}