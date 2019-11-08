// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudtrail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEventSelector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_resources": {
				Type: schema.TypeSet,
				Elem: propertyDataResource(),
				Required: false,
			},
			"include_management_events": {
				Type: schema.TypeBool,
				Required: false,
			},
			"read_write_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}