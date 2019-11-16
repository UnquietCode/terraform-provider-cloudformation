// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueDatabaseType string = "AWS::Glue::Database"

var glueDatabaseProperties map[string]string = map[string]string{
	"database_input": "DatabaseInput",
	"catalog_id": "CatalogId",
}

func ResourceGlueDatabase() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueDatabaseExists,
		Read: resourceGlueDatabaseRead,
		Create: resourceGlueDatabaseCreate,
		Update: resourceGlueDatabaseUpdate,
		Delete: resourceGlueDatabaseDelete,
		CustomizeDiff: resourceGlueDatabaseCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"database_input": {
				Type: schema.TypeList,
				Elem: propertyDatabaseDatabaseInput(),
				Required: true,
				MaxItems: 1,
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

func resourceGlueDatabaseExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueDatabaseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueDatabaseType, ResourceGlueDatabase(), data, meta)
}

func resourceGlueDatabaseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueDatabaseType, ResourceGlueDatabase(), data, glueDatabaseProperties, meta)
}

func resourceGlueDatabaseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueDatabaseType, ResourceGlueDatabase(), data, glueDatabaseProperties, meta)
}

func resourceGlueDatabaseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueDatabaseType, data, meta)
}

func resourceGlueDatabaseCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueDatabaseType, data, meta)
}
