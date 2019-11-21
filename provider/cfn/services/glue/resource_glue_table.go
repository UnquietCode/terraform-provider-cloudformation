// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueTableType string = "AWS::Glue::Table"

func ResourceGlueTable() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceGlueTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueTableType, ResourceGlueTable(), data, meta)
}

func resourceGlueTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueTableType, ResourceGlueTable(), data, meta)
}

func resourceGlueTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueTableType, ResourceGlueTable(), data, meta)
}

func resourceGlueTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueTableType, data, meta)
}

func resourceGlueTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueTableType, data, meta)
}
