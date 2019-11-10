// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationoutput.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsV2ApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisAnalyticsV2ApplicationOutputCreate,
		Read:   resourceKinesisAnalyticsV2ApplicationOutputRead,
		Update: resourceKinesisAnalyticsV2ApplicationOutputUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationOutputDelete,

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
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationOutputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}