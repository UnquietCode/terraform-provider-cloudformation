// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceGlueTable() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGlueTableRead,
		
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceGlueTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueTableType, DatasourceGlueTable(), data, meta)
}
