// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html

package iotevents

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTEventsDetectorModelType string = "AWS::IoTEvents::DetectorModel"

func DatasourceIoTEventsDetectorModel() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoTEventsDetectorModelRead,
		
		Schema: map[string]*schema.Schema{
			"detector_model_definition": {
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceIoTEventsDetectorModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTEventsDetectorModelType, DatasourceIoTEventsDetectorModel(), data, meta)
}
