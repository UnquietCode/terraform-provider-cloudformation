// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTargets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"s3_targets": {
				Type: schema.TypeList,
				Elem: propertyS3Target(),
				Required: false,
			},
			"catalog_targets": {
				Type: schema.TypeList,
				Elem: propertyCatalogTarget(),
				Required: false,
			},
			"jdbc_targets": {
				Type: schema.TypeList,
				Elem: propertyJdbcTarget(),
				Required: false,
			},
			"dynamo_db_targets": {
				Type: schema.TypeList,
				Elem: propertyDynamoDBTarget(),
				Required: false,
			},
		},
	}
}