// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeployment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"ignore_application_stop_failures": {
				Type: schema.TypeBool,
				Required: false,
			},
			"revision": {
				Type: schema.TypeList,
				Elem: propertyRevisionLocation(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}