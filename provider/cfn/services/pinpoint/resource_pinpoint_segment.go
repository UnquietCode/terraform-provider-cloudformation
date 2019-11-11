// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointSegment() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointSegmentCreate,
		Read:   resourcePinpointSegmentRead,
		Update: resourcePinpointSegmentUpdate,
		Delete: resourcePinpointSegmentDelete,

		Schema: map[string]*schema.Schema{
			"segment_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"segment_groups": {
				Type: schema.TypeList,
				Elem: propertySegmentGroups(),
				Optional: true,
				MaxItems: 1,
			},
			"dimensions": {
				Type: schema.TypeList,
				Elem: propertySegmentSegmentDimensions(),
				Optional: true,
				MaxItems: 1,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourcePinpointSegmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::Segment", ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::Segment", ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::Segment", ResourcePinpointSegment(), data, meta)
}

func resourcePinpointSegmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::Segment", data, meta)
}