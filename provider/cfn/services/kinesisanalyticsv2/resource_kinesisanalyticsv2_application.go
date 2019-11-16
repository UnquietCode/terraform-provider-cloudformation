// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationType string = "AWS::KinesisAnalyticsV2::Application"

var kinesisAnalyticsV2ApplicationProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"runtime_environment": "RuntimeEnvironment",
	"application_configuration": "ApplicationConfiguration",
	"application_description": "ApplicationDescription",
	"service_execution_role": "ServiceExecutionRole",
}

func ResourceKinesisAnalyticsV2Application() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisAnalyticsV2ApplicationExists,
		Read: resourceKinesisAnalyticsV2ApplicationRead,
		Create: resourceKinesisAnalyticsV2ApplicationCreate,
		Update: resourceKinesisAnalyticsV2ApplicationUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationDelete,
		CustomizeDiff: resourceKinesisAnalyticsV2ApplicationCustomizeDiff,
		
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
			},
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisAnalyticsV2ApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationType, ResourceKinesisAnalyticsV2Application(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsV2ApplicationType, ResourceKinesisAnalyticsV2Application(), data, kinesisAnalyticsV2ApplicationProperties, meta)
}

func resourceKinesisAnalyticsV2ApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsV2ApplicationType, ResourceKinesisAnalyticsV2Application(), data, kinesisAnalyticsV2ApplicationProperties, meta)
}

func resourceKinesisAnalyticsV2ApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsV2ApplicationType, data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsV2ApplicationType, data, meta)
}
