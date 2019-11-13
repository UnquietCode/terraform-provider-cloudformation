// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointEventStream() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEventStreamExists,
		Read:   resourcePinpointEventStreamRead,
		Create: resourcePinpointEventStreamCreate,
		Update: resourcePinpointEventStreamUpdate,
		Delete: resourcePinpointEventStreamDelete,
		CustomizeDiff: resourcePinpointEventStreamCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"destination_stream_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
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

func resourcePinpointEventStreamExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEventStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EventStream", ResourcePinpointEventStream(), data, meta)
}

func resourcePinpointEventStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EventStream", ResourcePinpointEventStream(), data, meta)
}

func resourcePinpointEventStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EventStream", ResourcePinpointEventStream(), data, meta)
}

func resourcePinpointEventStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EventStream", data, meta)
}

func resourcePinpointEventStreamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}