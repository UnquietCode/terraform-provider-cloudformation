// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointSegmentType string = "AWS::Pinpoint::Segment"

var pinpointSegmentProperties map[string]string = map[string]string{
	"segment_groups": "SegmentGroups",
	"dimensions": "Dimensions",
	"application_id": "ApplicationId",
	"tags": "Tags",
	"name": "Name",
}

func ResourcePinpointSegment() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointSegmentExists,
		Read: resourcePinpointSegmentRead,
		Create: resourcePinpointSegmentCreate,
		Update: resourcePinpointSegmentUpdate,
		Delete: resourcePinpointSegmentDelete,
		CustomizeDiff: resourcePinpointSegmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"segment_groups": {
				Type: schema.TypeSet,
				Elem: propertySegmentGroups(),
				Optional: true,
				MaxItems: 1,
			},
			"dimensions": {
				Type: schema.TypeSet,
				Elem: propertySegmentSegmentDimensions(),
				Optional: true,
				MaxItems: 1,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointSegmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointSegmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointSegmentType, ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointSegmentType, ResourcePinpointSegment(), data, pinpointSegmentProperties, meta)
}

func resourcePinpointSegmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointSegmentType, ResourcePinpointSegment(), data, pinpointSegmentProperties, meta)
}

func resourcePinpointSegmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointSegmentType, data, meta)
}

func resourcePinpointSegmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointSegmentType, data, meta)
}
