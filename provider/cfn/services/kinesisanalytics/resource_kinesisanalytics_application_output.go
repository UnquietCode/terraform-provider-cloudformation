// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisAnalyticsApplicationOutputCreate,
		Read:   resourceKinesisAnalyticsApplicationOutputRead,
		Update: resourceKinesisAnalyticsApplicationOutputUpdate,
		Delete: resourceKinesisAnalyticsApplicationOutputDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputOutput(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisAnalyticsApplicationOutputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalytics::ApplicationOutput", ResourceKinesisAnalyticsApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalytics::ApplicationOutput", ResourceKinesisAnalyticsApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsApplicationOutputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalytics::ApplicationOutput", ResourceKinesisAnalyticsApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsApplicationOutputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalytics::ApplicationOutput", data, meta)
}