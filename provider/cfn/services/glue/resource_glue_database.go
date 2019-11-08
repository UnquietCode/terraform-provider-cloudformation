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

func ResourceGlueDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueDatabaseCreate,
		Read:   resourceGlueDatabaseRead,
		Update: resourceGlueDatabaseUpdate,
		Delete: resourceGlueDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"database_input": {
				Type: schema.TypeList,
				Elem: propertyDatabaseInput(),
				Required: true,
				MaxItems: 1,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueDatabaseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Database", data, meta)
}

func resourceGlueDatabaseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Database", data, meta)
}

func resourceGlueDatabaseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Database", data, meta)
}

func resourceGlueDatabaseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Database", data, meta)
}