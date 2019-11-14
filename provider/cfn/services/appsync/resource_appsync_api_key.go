// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppSyncApiKey() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppSyncApiKeyExists,
		Read: resourceAppSyncApiKeyRead,
		Create: resourceAppSyncApiKeyCreate,
		Update: resourceAppSyncApiKeyUpdate,
		Delete: resourceAppSyncApiKeyDelete,
		CustomizeDiff: resourceAppSyncApiKeyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"expires": {
				Type: schema.TypeFloat,
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
			},
		},
	}
}

func resourceAppSyncApiKeyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppSyncApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::ApiKey", data, meta)
}

func resourceAppSyncApiKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

