// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTAnalyticsDatastore() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTAnalyticsDatastoreCreate,
		Read:   resourceIoTAnalyticsDatastoreRead,
		Update: resourceIoTAnalyticsDatastoreUpdate,
		Delete: resourceIoTAnalyticsDatastoreDelete,

		Schema: map[string]*schema.Schema{
			"datastore_storage": {
				Type: schema.TypeList,
				Elem: propertyDatastoreStorage(),
				Required: false,
				MaxItems: 1,
			},
			"datastore_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"retention_period": {
				Type: schema.TypeList,
				Elem: propertyRetentionPeriod(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceIoTAnalyticsDatastoreCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Datastore", data, meta)
}

func resourceIoTAnalyticsDatastoreRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTAnalytics::Datastore", data, meta)
}

func resourceIoTAnalyticsDatastoreUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Datastore", data, meta)
}

func resourceIoTAnalyticsDatastoreDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Datastore", data, meta)
}