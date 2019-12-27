// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html

package kinesisanalyticsv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationType string = "AWS::KinesisAnalyticsV2::Application"

func DatasourceKinesisAnalyticsV2Application() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKinesisAnalyticsV2ApplicationRead,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"runtime_environment": {
				Type: schema.TypeString,
				Required: true,
			},
			"application_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"application_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_execution_role": {
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

func datasourceKinesisAnalyticsV2ApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationType, DatasourceKinesisAnalyticsV2Application(), data, meta)
}
