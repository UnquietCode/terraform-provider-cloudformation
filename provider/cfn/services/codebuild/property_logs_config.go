// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLogsConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_watch_logs": {
				Type: schema.TypeList,
				Elem: propertyCloudWatchLogsConfig(),
				Required: false,
				MaxItems: 1,
			},
			"s3_logs": {
				Type: schema.TypeList,
				Elem: propertyS3LogsConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}