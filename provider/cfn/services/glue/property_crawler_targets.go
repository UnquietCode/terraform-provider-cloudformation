// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCrawlerTargets(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"s3_targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerS3Target(),
				Required: false,
			},
			"catalog_targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerCatalogTarget(),
				Required: false,
			},
			"jdbc_targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerJdbcTarget(),
				Required: false,
			},
			"dynamo_db_targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerDynamoDBTarget(),
				Required: false,
			},
		},
	}
}