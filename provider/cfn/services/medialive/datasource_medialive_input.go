// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html

package medialive

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaLiveInputType string = "AWS::MediaLive::Input"

func DatasourceMediaLiveInput() *schema.Resource {
	return &schema.Resource{
		Read: datasourceMediaLiveInputRead,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyInputInputDestinationRequest(),
				Optional: true,
			},
			"vpc": {
				Type: schema.TypeList,
				Elem: propertyInputInputVpcRequest(),
				Optional: true,
				MaxItems: 1,
			},
			"media_connect_flows": {
				Type: schema.TypeList,
				Elem: propertyInputMediaConnectFlowRequest(),
				Optional: true,
			},
			"input_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyInputInputSourceRequest(),
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
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

func datasourceMediaLiveInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaLiveInputType, DatasourceMediaLiveInput(), data, meta)
}
