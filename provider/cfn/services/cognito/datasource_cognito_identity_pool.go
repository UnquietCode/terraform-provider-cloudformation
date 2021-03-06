// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoIdentityPoolType string = "AWS::Cognito::IdentityPool"

func DatasourceCognitoIdentityPool() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCognitoIdentityPoolRead,
		
		Schema: map[string]*schema.Schema{
			"push_sync": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolPushSync(),
				Optional: true,
				MaxItems: 1,
			},
			"cognito_identity_providers": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolCognitoIdentityProvider(),
				Optional: true,
			},
			"cognito_events": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"developer_provider_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cognito_streams": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolCognitoStreams(),
				Optional: true,
				MaxItems: 1,
			},
			"identity_pool_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"allow_unauthenticated_identities": {
				Type: schema.TypeBool,
				Required: true,
			},
			"supported_login_providers": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"saml_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"open_id_connect_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allow_classic_flow": {
				Type: schema.TypeBool,
				Optional: true,
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

func datasourceCognitoIdentityPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoIdentityPoolType, DatasourceCognitoIdentityPool(), data, meta)
}
