// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"container_action": {
				Type: schema.TypeList,
				Elem: propertyContainerAction(),
				Required: false,
				MaxItems: 1,
			},
			"query_action": {
				Type: schema.TypeList,
				Elem: propertyQueryAction(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}