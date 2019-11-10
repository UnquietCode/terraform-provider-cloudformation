// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIdentity() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityCreate,
		Read:   resourceIdentityRead,
		Update: resourceIdentityUpdate,
		Delete: resourceIdentityDelete,

		Schema: map[string]*schema.Schema{
			"feedback_forwarding_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"dkim_signing_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyIdentityTags(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mail_from_attributes": {
				Type: schema.TypeList,
				Elem: propertyIdentityMailFromAttributes(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::Identity", data, meta)
}

func resourceIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::Identity", data, meta)
}

func resourceIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::Identity", data, meta)
}

func resourceIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::Identity", data, meta)
}