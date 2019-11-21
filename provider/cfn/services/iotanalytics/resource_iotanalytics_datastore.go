// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html

package iotanalytics

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTAnalyticsDatastoreType string = "AWS::IoTAnalytics::Datastore"

func ResourceIoTAnalyticsDatastore() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoTAnalyticsDatastoreRead,
		Create: resourceIoTAnalyticsDatastoreCreate,
		Update: resourceIoTAnalyticsDatastoreUpdate,
		Delete: resourceIoTAnalyticsDatastoreDelete,
		CustomizeDiff: resourceIoTAnalyticsDatastoreCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"datastore_storage": {
				Type: schema.TypeSet,
				Elem: propertyDatastoreDatastoreStorage(),
				Optional: true,
				MaxItems: 1,
			},
			"datastore_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retention_period": {
				Type: schema.TypeSet,
				Elem: propertyDatastoreRetentionPeriod(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceIoTAnalyticsDatastoreRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsDatastoreType, ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTAnalyticsDatastoreType, ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTAnalyticsDatastoreType, ResourceIoTAnalyticsDatastore(), data, meta)
}

func resourceIoTAnalyticsDatastoreDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTAnalyticsDatastoreType, data, meta)
}

func resourceIoTAnalyticsDatastoreCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTAnalyticsDatastoreType, data, meta)
}
