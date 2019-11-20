// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceKinesisAnalyticsApplication() *schema.Resource {
	return &schema.Resource{
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

func resourceKinesisAnalyticsApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, meta)
}

func resourceKinesisAnalyticsApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, meta)
}

func resourceKinesisAnalyticsApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsApplicationType, ResourceKinesisAnalyticsApplication(), data, meta)
}

func resourceKinesisAnalyticsApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsApplicationType, data, meta)
}

func resourceKinesisAnalyticsApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsApplicationType, data, meta)
}
