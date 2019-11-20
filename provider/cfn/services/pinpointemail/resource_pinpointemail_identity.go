// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailIdentityType string = "AWS::PinpointEmail::Identity"

func ResourcePinpointEmailIdentity() *schema.Resource {
	return &schema.Resource{
		Read: resourcePinpointEmailIdentityRead,
		Create: resourcePinpointEmailIdentityCreate,
		Update: resourcePinpointEmailIdentityUpdate,
		Delete: resourcePinpointEmailIdentityDelete,
		CustomizeDiff: resourcePinpointEmailIdentityCustomizeDiff,
		
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
				Type: schema.TypeSet,
				Elem: propertyIdentityMailFromAttributes(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointEmailIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailIdentityType, ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointEmailIdentityType, ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointEmailIdentityType, ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointEmailIdentityType, data, meta)
}

func resourcePinpointEmailIdentityCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointEmailIdentityType, data, meta)
}
