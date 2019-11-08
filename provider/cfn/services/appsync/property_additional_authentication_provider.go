// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAdditionalAuthenticationProvider() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"open_id_connect_config": {
				Type: schema.TypeList,
				Elem: propertyOpenIDConnectConfig(),
				Required: false,
				MaxItems: 1,
			},
			"user_pool_config": {
				Type: schema.TypeList,
				Elem: propertyCognitoUserPoolConfig(),
				Required: false,
				MaxItems: 1,
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}