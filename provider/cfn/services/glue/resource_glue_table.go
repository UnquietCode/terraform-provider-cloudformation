// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueTableCreate,
		Read:   resourceGlueTableRead,
		Update: resourceGlueTableUpdate,
		Delete: resourceGlueTableDelete,

		Schema: map[string]*schema.Schema{
			"table_input": {
				Type: schema.TypeList,
				Elem: propertyTableInput(),
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

func resourceGlueTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Table", data, meta)
}

func resourceGlueTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Table", data, meta)
}

func resourceGlueTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Table", data, meta)
}

func resourceGlueTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Table", data, meta)
}