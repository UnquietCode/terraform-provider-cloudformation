// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-schemaconfiguration.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var deliveryStreamSchemaConfigurationProperties map[string]string = map[string]string{
	"catalog_id": "CatalogId",
	"database_name": "DatabaseName",
	"region": "Region",
	"role_arn": "RoleARN",
	"table_name": "TableName",
	"version_id": "VersionId",
}

func propertyDeliveryStreamSchemaConfiguration(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"region": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"version_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
