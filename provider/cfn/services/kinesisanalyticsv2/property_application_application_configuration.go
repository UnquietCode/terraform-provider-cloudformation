// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationApplicationConfiguration(extras...string) *schema.Resource {
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