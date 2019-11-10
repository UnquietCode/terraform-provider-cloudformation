// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAuthorizer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthorizerCreate,
		Read:   resourceAuthorizerRead,
		Update: resourceAuthorizerUpdate,
		Delete: resourceAuthorizerDelete,

		Schema: map[string]*schema.Schema{
			"auth_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_credentials": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_result_ttl_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"authorizer_uri": {
				Type: schema.TypeString,
				Required: false,
			},
			"identity_source": {
				Type: schema.TypeString,
				Required: false,
			},
			"identity_validation_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"provider_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAuthorizerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Authorizer", data, meta)
}

func resourceAuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Authorizer", data, meta)
}

func resourceAuthorizerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Authorizer", data, meta)
}

func resourceAuthorizerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Authorizer", data, meta)
}