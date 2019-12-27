// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceAppSyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppSyncGraphQLApiRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAppSyncGraphQLApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncGraphQLApiType, DatasourceAppSyncGraphQLApi(), data, meta)
}
