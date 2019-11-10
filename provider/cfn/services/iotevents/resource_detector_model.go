// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDetectorModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceDetectorModelCreate,
		Read:   resourceDetectorModelRead,
		Update: resourceDetectorModelUpdate,
		Delete: resourceDetectorModelDelete,

		Schema: map[string]*schema.Schema{
			"detector_model_definition": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelDetectorModelDefinition(),
				Required: false,
				MaxItems: 1,
			},
			"detector_model_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"detector_model_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"key": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDetectorModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTEvents::DetectorModel", data, meta)
}

func resourceDetectorModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTEvents::DetectorModel", data, meta)
}

func resourceDetectorModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTEvents::DetectorModel", data, meta)
}

func resourceDetectorModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTEvents::DetectorModel", data, meta)
}