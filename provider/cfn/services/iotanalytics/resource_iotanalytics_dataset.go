// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const ioTAnalyticsDatasetType string = "AWS::IoTAnalytics::Dataset"

var ioTAnalyticsDatasetProperties map[string]string = map[string]string{
	"actions": "Actions",
	"dataset_name": "DatasetName",
	"content_delivery_rules": "ContentDeliveryRules",
	"triggers": "Triggers",
	"versioning_configuration": "VersioningConfiguration",
	"retention_period": "RetentionPeriod",
	"tags": "Tags",
}

func ResourceIoTAnalyticsDataset() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTAnalyticsDatasetExists,
		Read: resourceIoTAnalyticsDatasetRead,
		Create: resourceIoTAnalyticsDatasetCreate,
		Update: resourceIoTAnalyticsDatasetUpdate,
		Delete: resourceIoTAnalyticsDatasetDelete,
		CustomizeDiff: resourceIoTAnalyticsDatasetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeList,
				Elem: propertyDatasetAction(),
				Required: true,
			},
			"dataset_name": {
				Type: schema.TypeString,
				Optional: true,
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
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTAnalyticsDatasetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTAnalyticsDatasetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsDatasetType, ResourceIoTAnalyticsDataset(), data, meta)
}

func resourceIoTAnalyticsDatasetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTAnalyticsDatasetType, ResourceIoTAnalyticsDataset(), data, ioTAnalyticsDatasetProperties, meta)
}

func resourceIoTAnalyticsDatasetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTAnalyticsDatasetType, ResourceIoTAnalyticsDataset(), data, ioTAnalyticsDatasetProperties, meta)
}

func resourceIoTAnalyticsDatasetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTAnalyticsDatasetType, data, meta)
}

func resourceIoTAnalyticsDatasetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTAnalyticsDatasetType, data, meta)
}
