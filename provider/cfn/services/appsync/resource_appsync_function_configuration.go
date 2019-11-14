// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-functionconfiguration.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppSyncFunctionConfiguration() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppSyncFunctionConfigurationExists,
		Read: resourceAppSyncFunctionConfigurationRead,
		Create: resourceAppSyncFunctionConfigurationCreate,
		Update: resourceAppSyncFunctionConfigurationUpdate,
		Delete: resourceAppSyncFunctionConfigurationDelete,
		CustomizeDiff: resourceAppSyncFunctionConfigurationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"data_source_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_mapping_template": {
				Type: schema.TypeString,
				Optional: true,
			},
			"response_mapping_template": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func resourceAppSyncFunctionConfigurationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppSyncFunctionConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::FunctionConfiguration", data, meta)
}

func resourceAppSyncFunctionConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

