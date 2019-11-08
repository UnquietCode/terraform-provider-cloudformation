// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scan_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"recipients": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: propertyAction(),
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tls_policy": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}