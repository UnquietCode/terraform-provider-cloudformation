// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncApiKeyType string = "AWS::AppSync::ApiKey"

func ResourceAppSyncApiKey() *schema.Resource {
	return &schema.Resource{
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

func resourceAppSyncApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncApiKeyType, ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncApiKeyType, ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncApiKeyType, ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncApiKeyType, data, meta)
}

func resourceAppSyncApiKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncApiKeyType, data, meta)
}
