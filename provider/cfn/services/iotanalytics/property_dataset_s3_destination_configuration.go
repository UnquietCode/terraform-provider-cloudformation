// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html

package iotanalytics

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var datasetS3DestinationConfigurationProperties map[string]string = map[string]string{
	"glue_configuration": "GlueConfiguration",
	"bucket": "Bucket",
	"key": "Key",
	"role_arn": "RoleArn",
}

func propertyDatasetS3DestinationConfiguration(extras...string) *schema.Resource {
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
			"glue_configuration": {
				Type: schema.TypeList,
				Elem: propertyDatasetGlueConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"key": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
