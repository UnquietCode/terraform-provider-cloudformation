// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTAnalyticsDatastore() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTAnalyticsDatastoreExists,
		Read:   resourceIoTAnalyticsDatastoreRead,
		Create: resourceIoTAnalyticsDatastoreCreate,
		Update: resourceIoTAnalyticsDatastoreUpdate,
		Delete: resourceIoTAnalyticsDatastoreDelete,
		CustomizeDiff: resourceIoTAnalyticsDatastoreCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"datastore_storage": {
				Type: schema.TypeList,
				Elem: propertyDatastoreDatastoreStorage(),
				Optional: true,
				MaxItems: 1,
			},
			"datastore_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retention_period": {
				Type: schema.TypeList,
				Elem: propertyDatastoreRetentionPeriod(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceIoTAnalyticsDatastoreExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTAnalyticsDatastoreRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTAnalytics::Datastore", ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Datastore", ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Datastore", ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Datastore", data, meta)
}

func resourceIoTAnalyticsDatastoreCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}