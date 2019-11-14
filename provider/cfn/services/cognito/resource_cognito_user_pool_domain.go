// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolDomain() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolDomainExists,
		Read: resourceCognitoUserPoolDomainRead,
		Create: resourceCognitoUserPoolDomainCreate,
		Update: resourceCognitoUserPoolDomainUpdate,
		Delete: resourceCognitoUserPoolDomainDelete,
		CustomizeDiff: resourceCognitoUserPoolDomainCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"custom_domain_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolDomainCustomDomainConfigType(),
				Optional: true,
				MaxItems: 1,
			},
			"domain": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCognitoUserPoolDomainExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolDomain", ResourceCognitoUserPoolDomain(), data, meta)
}

func resourceCognitoUserPoolDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolDomain", ResourceCognitoUserPoolDomain(), data, meta)
}

func resourceCognitoUserPoolDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolDomain", ResourceCognitoUserPoolDomain(), data, meta)
}

func resourceCognitoUserPoolDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolDomain", data, meta)
}

func resourceCognitoUserPoolDomainCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

