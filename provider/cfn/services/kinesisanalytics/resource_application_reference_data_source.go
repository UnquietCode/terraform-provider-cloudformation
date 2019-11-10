// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplicationReferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationReferenceDataSourceCreate,
		Read:   resourceApplicationReferenceDataSourceRead,
		Update: resourceApplicationReferenceDataSourceUpdate,
		Delete: resourceApplicationReferenceDataSourceDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"reference_data_source": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceReferenceDataSource(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceApplicationReferenceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalytics::ApplicationReferenceDataSource", data, meta)
}

func resourceApplicationReferenceDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalytics::ApplicationReferenceDataSource", data, meta)
}

func resourceApplicationReferenceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalytics::ApplicationReferenceDataSource", data, meta)
}

func resourceApplicationReferenceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalytics::ApplicationReferenceDataSource", data, meta)
}