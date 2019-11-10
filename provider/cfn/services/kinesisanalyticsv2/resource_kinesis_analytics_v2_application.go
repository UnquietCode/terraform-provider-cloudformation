// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsV2Application() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisAnalyticsV2ApplicationCreate,
		Read:   resourceKinesisAnalyticsV2ApplicationRead,
		Update: resourceKinesisAnalyticsV2ApplicationUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"runtime_environment": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"application_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_execution_role": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::Application", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::Application", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::Application", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::Application", data, meta)
}