// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html

package pinpoint

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointSegmentType string = "AWS::Pinpoint::Segment"

func ResourcePinpointSegment() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourcePinpointSegmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointSegmentType, ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointSegmentType, ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointSegmentType, ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointSegmentType, data, meta)
}

func resourcePinpointSegmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointSegmentType, data, meta)
}
