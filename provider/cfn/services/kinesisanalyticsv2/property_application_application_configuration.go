// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_code_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationApplicationCodeConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"environment_properties": {
				Type: schema.TypeSet,
				Elem: propertyApplicationEnvironmentProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"flink_application_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationFlinkApplicationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"sql_application_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationSqlApplicationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"application_snapshot_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationApplicationSnapshotConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
