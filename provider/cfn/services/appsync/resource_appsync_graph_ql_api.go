// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html

package appsync

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncGraphQLApiType string = "AWS::AppSync::GraphQLApi"

func ResourceAppSyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppSyncGraphQLApiRead,
		Create: resourceAppSyncGraphQLApiCreate,
		Update: resourceAppSyncGraphQLApiUpdate,
		Delete: resourceAppSyncGraphQLApiDelete,
		CustomizeDiff: resourceAppSyncGraphQLApiCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"open_id_connect_config": {
				Type: schema.TypeSet,
				Elem: propertyGraphQLApiOpenIDConnectConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"user_pool_config": {
				Type: schema.TypeSet,
				Elem: propertyGraphQLApiUserPoolConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyGraphQLApiLogConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"additional_authentication_providers": {
				Type: schema.TypeSet,
				Elem: propertyGraphQLApiAdditionalAuthenticationProviders(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAppSyncGraphQLApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncGraphQLApiType, ResourceAppSyncGraphQLApi(), data, meta)
}

func resourceAppSyncGraphQLApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncGraphQLApiType, data, meta)
}

func resourceAppSyncGraphQLApiCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncGraphQLApiType, data, meta)
}
