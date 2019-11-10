// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceAppSyncApiKeyCreate,
		Read:   resourceAppSyncApiKeyRead,
		Update: resourceAppSyncApiKeyUpdate,
		Delete: resourceAppSyncApiKeyDelete,

		Schema: map[string]*schema.Schema{
			"api_key": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				ForceNew: true,
			},
		},
	}
}

func resourceAppSyncApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::ApiKey", ResourceAppSyncApiKey(), data, meta)
}

func resourceAppSyncApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::ApiKey", data, meta)
}