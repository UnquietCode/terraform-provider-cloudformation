// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html

package pinpoint

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailChannelType string = "AWS::Pinpoint::EmailChannel"

func DatasourcePinpointEmailChannel() *schema.Resource {
	return &schema.Resource{
		Read: datasourcePinpointEmailChannelRead,
		
		Schema: map[string]*schema.Schema{
			"configuration_set": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_address": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"identity": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
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

func datasourcePinpointEmailChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailChannelType, DatasourcePinpointEmailChannel(), data, meta)
}
