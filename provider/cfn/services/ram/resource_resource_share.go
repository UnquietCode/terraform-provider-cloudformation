// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ram

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResourceShare() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourceShareCreate,
		Read:   resourceResourceShareRead,
		Update: resourceResourceShareUpdate,
		Delete: resourceResourceShareDelete,

		Schema: map[string]*schema.Schema{
			"principals": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"allow_external_principals": {
				Type: schema.TypeBool,
				Required: false,
			},
			"resource_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceResourceShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RAM::ResourceShare", data, meta)
}

func resourceResourceShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RAM::ResourceShare", data, meta)
}

func resourceResourceShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RAM::ResourceShare", data, meta)
}

func resourceResourceShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RAM::ResourceShare", data, meta)
}