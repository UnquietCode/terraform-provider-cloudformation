// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsApplicationOutputType string = "AWS::KinesisAnalytics::ApplicationOutput"

var kinesisAnalyticsApplicationOutputProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"output": "Output",
}

func ResourceKinesisAnalyticsApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisAnalyticsApplicationOutputExists,
		Read: resourceKinesisAnalyticsApplicationOutputRead,
		Create: resourceKinesisAnalyticsApplicationOutputCreate,
		Update: resourceKinesisAnalyticsApplicationOutputUpdate,
		Delete: resourceKinesisAnalyticsApplicationOutputDelete,
		CustomizeDiff: resourceKinesisAnalyticsApplicationOutputCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"output": {
				Type: schema.TypeSet,
				Elem: propertyApplicationOutputOutput(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisAnalyticsApplicationOutputExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisAnalyticsApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsApplicationOutputType, ResourceKinesisAnalyticsApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsApplicationOutputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsApplicationOutputType, ResourceKinesisAnalyticsApplicationOutput(), data, kinesisAnalyticsApplicationOutputProperties, meta)
}

func resourceKinesisAnalyticsApplicationOutputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsApplicationOutputType, ResourceKinesisAnalyticsApplicationOutput(), data, kinesisAnalyticsApplicationOutputProperties, meta)
}

func resourceKinesisAnalyticsApplicationOutputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsApplicationOutputType, data, meta)
}

func resourceKinesisAnalyticsApplicationOutputCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsApplicationOutputType, data, meta)
}
