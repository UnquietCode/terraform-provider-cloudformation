// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorsession.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TrafficMirrorSessionType string = "AWS::EC2::TrafficMirrorSession"

func DatasourceEC2TrafficMirrorSession() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2TrafficMirrorSessionRead,
		
		Schema: map[string]*schema.Schema{
			"traffic_mirror_target_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"session_number": {
				Type: schema.TypeInt,
				Required: true,
			},
			"virtual_network_id": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"packet_length": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"traffic_mirror_filter_id": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceEC2TrafficMirrorSessionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TrafficMirrorSessionType, DatasourceEC2TrafficMirrorSession(), data, meta)
}
