// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationApplicationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_code_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationCodeConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"environment_properties": {
				Type: schema.TypeList,
				Elem: propertyApplicationEnvironmentProperties(),
				Required: false,
				MaxItems: 1,
			},
			"flink_application_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationFlinkApplicationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"sql_application_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationSqlApplicationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"application_snapshot_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationSnapshotConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}