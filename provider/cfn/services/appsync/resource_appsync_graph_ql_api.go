// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncGraphQLApiType string = "AWS::AppSync::GraphQLApi"

var appSyncGraphQLApiProperties map[string]string = map[string]string{
	"open_id_connect_config": "OpenIDConnectConfig",
	"user_pool_config": "UserPoolConfig",
	"tags": "Tags",
	"name": "Name",
	"authentication_type": "AuthenticationType",
	"log_config": "LogConfig",
	"additional_authentication_providers": "AdditionalAuthenticationProviders",
}

func ResourceAppSyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppSyncGraphQLApiExists,
		Read: resourceAppSyncGraphQLApiRead,
		Create: resourceAppSyncGraphQLApiCreate,
		Update: resourceAppSyncGraphQLApiUpdate,
		Delete: resourceAppSyncGraphQLApiDelete,
		CustomizeDiff: resourceAppSyncGraphQLApiCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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

func resourceAppSyncGraphQLApiExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppSyncGraphQLApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, appSyncGraphQLApiProperties, meta)
}

func resourceAppSyncGraphQLApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, appSyncGraphQLApiProperties, meta)
}

func resourceAppSyncGraphQLApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncGraphQLApiType, data, meta)
}

func resourceAppSyncGraphQLApiCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncGraphQLApiType, data, meta)
}
