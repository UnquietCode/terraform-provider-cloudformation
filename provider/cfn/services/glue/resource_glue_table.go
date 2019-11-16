// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueTableType string = "AWS::Glue::Table"

var glueTableProperties map[string]string = map[string]string{
	"table_input": "TableInput",
	"database_name": "DatabaseName",
	"catalog_id": "CatalogId",
}

func ResourceGlueTable() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueTableExists,
		Read: resourceGlueTableRead,
		Create: resourceGlueTableCreate,
		Update: resourceGlueTableUpdate,
		Delete: resourceGlueTableDelete,
		CustomizeDiff: resourceGlueTableCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"table_input": {
				Type: schema.TypeSet,
				Elem: propertyTableTableInput(),
				Required: true,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueTableExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueTableType, ResourceGlueTable(), data, meta)
}

func resourceGlueTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueTableType, ResourceGlueTable(), data, glueTableProperties, meta)
}

func resourceGlueTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueTableType, ResourceGlueTable(), data, glueTableProperties, meta)
}

func resourceGlueTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueTableType, data, meta)
}

func resourceGlueTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueTableType, data, meta)
}
