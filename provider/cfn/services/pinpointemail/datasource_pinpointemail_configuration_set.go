// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html

package pinpointemail

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailConfigurationSetType string = "AWS::PinpointEmail::ConfigurationSet"

func DatasourcePinpointEmailConfigurationSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourcePinpointEmailConfigurationSetRead,
		
		Schema: map[string]*schema.Schema{
			"sending_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetSendingOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tracking_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTrackingOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"reputation_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetReputationOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"delivery_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetDeliveryOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTags(),
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourcePinpointEmailConfigurationSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailConfigurationSetType, DatasourcePinpointEmailConfigurationSet(), data, meta)
}
