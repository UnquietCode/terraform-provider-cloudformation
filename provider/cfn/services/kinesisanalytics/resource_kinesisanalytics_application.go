// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsApplicationType string = "AWS::KinesisAnalytics::Application"

var kinesisAnalyticsApplicationProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"inputs": "Inputs",
	"application_description": "ApplicationDescription",
	"application_code": "ApplicationCode",
}

func ResourceKinesisAnalyticsApplication() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisAnalyticsApplicationExists,
		Read: resourceKinesisAnalyticsApplicationRead,
		Create: resourceKinesisAnalyticsApplicationCreate,
		Update: resourceKinesisAnalyticsApplicationUpdate,
		Delete: resourceKinesisAnalyticsApplicationDelete,
		CustomizeDiff: resourceKinesisAnalyticsApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"inputs": {
				Type: schema.TypeList,
				Elem: propertyApplicationInput(),
				Required: true,
			},
			"application_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application_code": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisAnalyticsApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisAnalyticsApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, meta)
}

func resourceKinesisAnalyticsApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, kinesisAnalyticsApplicationProperties, meta)
}

func resourceKinesisAnalyticsApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, kinesisAnalyticsApplicationProperties, meta)
}

func resourceKinesisAnalyticsApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsApplicationType, data, meta)
}

func resourceKinesisAnalyticsApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsApplicationType, data, meta)
}
