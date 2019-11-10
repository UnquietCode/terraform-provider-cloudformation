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

func ResourceAppSyncGraphQLSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppSyncGraphQLSchemaCreate,
		Read:   resourceAppSyncGraphQLSchemaRead,
		Update: resourceAppSyncGraphQLSchemaUpdate,
		Delete: resourceAppSyncGraphQLSchemaDelete,

		Schema: map[string]*schema.Schema{
			"definition": {
				Type: schema.TypeString,
				Required: false,
			},
			"definition_s3_location": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppSyncGraphQLSchemaCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::GraphQLSchema", data, meta)
}

func resourceAppSyncGraphQLSchemaRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::GraphQLSchema", data, meta)
}

func resourceAppSyncGraphQLSchemaUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::GraphQLSchema", data, meta)
}

func resourceAppSyncGraphQLSchemaDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::GraphQLSchema", data, meta)
}