// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventStreamCreate,
		Read:   resourceEventStreamRead,
		Update: resourceEventStreamUpdate,
		Delete: resourceEventStreamDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_stream_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEventStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EventStream", data, meta)
}

func resourceEventStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EventStream", data, meta)
}

func resourceEventStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EventStream", data, meta)
}

func resourceEventStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EventStream", data, meta)
}