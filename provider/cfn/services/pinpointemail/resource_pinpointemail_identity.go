// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointEmailIdentity() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEmailIdentityExists,
		Read:   resourcePinpointEmailIdentityRead,
		Create: resourcePinpointEmailIdentityCreate,
		Update: resourcePinpointEmailIdentityUpdate,
		Delete: resourcePinpointEmailIdentityDelete,
		
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
			},
		},
	}
}

func resourcePinpointEmailIdentityExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::Identity", ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::Identity", ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::Identity", ResourcePinpointEmailIdentity(), data, meta)
}

func resourcePinpointEmailIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::Identity", data, meta)
}