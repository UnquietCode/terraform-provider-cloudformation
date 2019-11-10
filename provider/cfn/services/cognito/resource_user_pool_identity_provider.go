// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceUserPoolIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPoolIdentityProviderCreate,
		Read:   resourceUserPoolIdentityProviderRead,
		Update: resourceUserPoolIdentityProviderUpdate,
		Delete: resourceUserPoolIdentityProviderDelete,

		Schema: map[string]*schema.Schema{
			"provider_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attribute_mapping": {
				Type: schema.TypeMap,
				Required: false,
			},
			"provider_details": {
				Type: schema.TypeMap,
				Required: false,
			},
			"provider_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"idp_identifiers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceUserPoolIdentityProviderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolIdentityProvider", data, meta)
}

func resourceUserPoolIdentityProviderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolIdentityProvider", data, meta)
}

func resourceUserPoolIdentityProviderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolIdentityProvider", data, meta)
}

func resourceUserPoolIdentityProviderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolIdentityProvider", data, meta)
}