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

func ResourceUserPoolDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPoolDomainCreate,
		Read:   resourceUserPoolDomainRead,
		Update: resourceUserPoolDomainUpdate,
		Delete: resourceUserPoolDomainDelete,

		Schema: map[string]*schema.Schema{
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"custom_domain_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolDomainCustomDomainConfigType(),
				Required: false,
				MaxItems: 1,
			},
			"domain": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceUserPoolDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolDomain", data, meta)
}

func resourceUserPoolDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolDomain", data, meta)
}

func resourceUserPoolDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolDomain", data, meta)
}

func resourceUserPoolDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolDomain", data, meta)
}