// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html

package appsync

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncGraphQLSchemaType string = "AWS::AppSync::GraphQLSchema"

func ResourceAppSyncGraphQLSchema() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppSyncGraphQLSchemaRead,
		Create: resourceAppSyncGraphQLSchemaCreate,
		Update: resourceAppSyncGraphQLSchemaUpdate,
		Delete: resourceAppSyncGraphQLSchemaDelete,
		CustomizeDiff: resourceAppSyncGraphQLSchemaCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"definition": {
				Type: schema.TypeString,
				Optional: true,
			},
			"definition_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
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

func resourceAppSyncGraphQLSchemaRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncGraphQLSchemaType, ResourceAppSyncGraphQLSchema(), data, meta)
}

func resourceAppSyncGraphQLSchemaCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncGraphQLSchemaType, ResourceAppSyncGraphQLSchema(), data, meta)
}

func resourceAppSyncGraphQLSchemaUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncGraphQLSchemaType, ResourceAppSyncGraphQLSchema(), data, meta)
}

func resourceAppSyncGraphQLSchemaDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncGraphQLSchemaType, data, meta)
}

func resourceAppSyncGraphQLSchemaCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncGraphQLSchemaType, data, meta)
}
