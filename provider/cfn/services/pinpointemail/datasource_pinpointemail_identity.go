// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html

package pinpointemail

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailIdentityType string = "AWS::PinpointEmail::Identity"

func DatasourcePinpointEmailIdentity() *schema.Resource {
	return &schema.Resource{
		Read: datasourcePinpointEmailIdentityRead,
		
		Schema: map[string]*schema.Schema{
			"feedback_forwarding_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"dkim_signing_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyIdentityTags(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"mail_from_attributes": {
				Type: schema.TypeList,
				Elem: propertyIdentityMailFromAttributes(),
				Optional: true,
				MaxItems: 1,
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

func datasourcePinpointEmailIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailIdentityType, DatasourcePinpointEmailIdentity(), data, meta)
}
