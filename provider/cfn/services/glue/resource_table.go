// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceTableCreate,
		Read:   resourceTableRead,
		Update: resourceTableUpdate,
		Delete: resourceTableDelete,

		Schema: map[string]*schema.Schema{
			"table_input": {
				Type: schema.TypeList,
				Elem: propertyTableTableInput(),
				Required: true,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Table", data, meta)
}

func resourceTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Table", data, meta)
}

func resourceTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Table", data, meta)
}

func resourceTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Table", data, meta)
}