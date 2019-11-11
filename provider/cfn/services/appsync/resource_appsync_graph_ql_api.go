// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppSyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppSyncGraphQLApiCreate,
		Read:   resourceAppSyncGraphQLApiRead,
		Update: resourceAppSyncGraphQLApiUpdate,
		Delete: resourceAppSyncGraphQLApiDelete,

		Schema: map[string]*schema.Schema{
			"graph_ql_url": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"open_id_connect_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiOpenIDConnectConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"user_pool_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiUserPoolConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiTags(),
				Optional: true,
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
				Optional: true,
				MaxItems: 1,
			},
			"additional_authentication_providers": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiAdditionalAuthenticationProviders(),
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

func resourceAppSyncGraphQLApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::GraphQLApi", ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::GraphQLApi", ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::GraphQLApi", ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::GraphQLApi", data, meta)
}