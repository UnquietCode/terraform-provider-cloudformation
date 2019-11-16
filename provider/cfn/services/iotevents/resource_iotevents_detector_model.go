// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTEventsDetectorModelType string = "AWS::IoTEvents::DetectorModel"

var ioTEventsDetectorModelProperties map[string]string = map[string]string{
	"detector_model_definition": "DetectorModelDefinition",
	"detector_model_name": "DetectorModelName",
	"detector_model_description": "DetectorModelDescription",
	"key": "Key",
	"role_arn": "RoleArn",
	"tags": "Tags",
}

func ResourceIoTEventsDetectorModel() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTEventsDetectorModelExists,
		Read: resourceIoTEventsDetectorModelRead,
		Create: resourceIoTEventsDetectorModelCreate,
		Update: resourceIoTEventsDetectorModelUpdate,
		Delete: resourceIoTEventsDetectorModelDelete,
		CustomizeDiff: resourceIoTEventsDetectorModelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"detector_model_definition": {
				Type: schema.TypeSet,
				Elem: propertyDetectorModelDetectorModelDefinition(),
				Optional: true,
				MaxItems: 1,
			},
			"detector_model_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"detector_model_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceIoTEventsDetectorModelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTEventsDetectorModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTEventsDetectorModelType, ResourceIoTEventsDetectorModel(), data, meta)
}

func resourceIoTEventsDetectorModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTEventsDetectorModelType, ResourceIoTEventsDetectorModel(), data, ioTEventsDetectorModelProperties, meta)
}

func resourceIoTEventsDetectorModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTEventsDetectorModelType, ResourceIoTEventsDetectorModel(), data, ioTEventsDetectorModelProperties, meta)
}

func resourceIoTEventsDetectorModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTEventsDetectorModelType, data, meta)
}

func resourceIoTEventsDetectorModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTEventsDetectorModelType, data, meta)
}
