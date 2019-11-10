// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-dataset.html

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTAnalyticsDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTAnalyticsDatasetCreate,
		Read:   resourceIoTAnalyticsDatasetRead,
		Update: resourceIoTAnalyticsDatasetUpdate,
		Delete: resourceIoTAnalyticsDatasetDelete,

		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeList,
				Elem: propertyDatasetAction(),
				Required: true,
			},
			"dataset_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content_delivery_rules": {
				Type: schema.TypeList,
				Elem: propertyDatasetDatasetContentDeliveryRule(),
				Optional: true,
			},
			"triggers": {
				Type: schema.TypeList,
				Elem: propertyDatasetTrigger(),
				Optional: true,
			},
			"versioning_configuration": {
				Type: schema.TypeList,
				Elem: propertyDatasetVersioningConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"retention_period": {
				Type: schema.TypeList,
				Elem: propertyDatasetRetentionPeriod(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceIoTAnalyticsDatasetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Dataset", data, meta)
}

func resourceIoTAnalyticsDatasetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTAnalytics::Dataset", data, meta)
}

func resourceIoTAnalyticsDatasetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Dataset", data, meta)
}

func resourceIoTAnalyticsDatasetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Dataset", data, meta)
}