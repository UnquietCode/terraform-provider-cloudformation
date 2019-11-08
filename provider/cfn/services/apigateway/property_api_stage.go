// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApiStage() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"stage": {
				Type: schema.TypeString,
				Required: false,
			},
			"throttle": {
				Type: schema.TypeMap,
				Elem: propertyThrottleSettings(),
				Required: false,
			},
		},
	}
}