// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"expires": {
				Type: schema.TypeFloat,
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

func resourceAppSyncApiKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::ApiKey", data, meta)
}

func resourceAppSyncApiKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::ApiKey", data, meta)
}

func resourceAppSyncApiKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::ApiKey", data, meta)
}

func resourceAppSyncApiKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::ApiKey", data, meta)
}