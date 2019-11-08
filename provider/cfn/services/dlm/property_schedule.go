// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dlm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

func propertySchedule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tags_to_add": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"create_rule": {
				Type: schema.TypeList,
				Elem: propertyCreateRule(),
				Required: false,
				MaxItems: 1,
			},
			"variable_tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"fast_restore_rule": {
				Type: schema.TypeList,
				Elem: propertyFastRestoreRule(),
				Required: false,
				MaxItems: 1,
			},
			"retain_rule": {
				Type: schema.TypeList,
				Elem: propertyRetainRule(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"copy_tags": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}