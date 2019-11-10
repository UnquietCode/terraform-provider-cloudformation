// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraphQLApiCreate,
		Read:   resourceGraphQLApiRead,
		Update: resourceGraphQLApiUpdate,
		Delete: resourceGraphQLApiDelete,

		Schema: map[string]*schema.Schema{
			"open_id_connect_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiOpenIDConnectConfig(),
				Required: false,
				MaxItems: 1,
			},
			"user_pool_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiUserPoolConfig(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiTags(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiLogConfig(),
				Required: false,
				MaxItems: 1,
			},
			"additional_authentication_providers": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiAdditionalAuthenticationProviders(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceGraphQLApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::GraphQLApi", data, meta)
}

func resourceGraphQLApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::GraphQLApi", data, meta)
}

func resourceGraphQLApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::GraphQLApi", data, meta)
}

func resourceGraphQLApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::GraphQLApi", data, meta)
}