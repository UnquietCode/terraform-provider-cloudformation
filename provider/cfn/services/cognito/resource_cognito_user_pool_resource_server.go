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

func ResourceCognitoUserPoolResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCognitoUserPoolResourceServerCreate,
		Read:   resourceCognitoUserPoolResourceServerRead,
		Update: resourceCognitoUserPoolResourceServerUpdate,
		Delete: resourceCognitoUserPoolResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identifier": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scopes": {
				Type: schema.TypeList,
				Elem: propertyUserPoolResourceServerResourceServerScopeType(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCognitoUserPoolResourceServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolResourceServer", data, meta)
}

func resourceCognitoUserPoolResourceServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolResourceServer", data, meta)
}

func resourceCognitoUserPoolResourceServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolResourceServer", data, meta)
}

func resourceCognitoUserPoolResourceServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolResourceServer", data, meta)
}