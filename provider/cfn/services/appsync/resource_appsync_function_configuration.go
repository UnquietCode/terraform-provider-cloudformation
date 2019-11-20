// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-functionconfiguration.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncFunctionConfigurationType string = "AWS::AppSync::FunctionConfiguration"

func ResourceAppSyncFunctionConfiguration() *schema.Resource {
	return &schema.Resource{
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

func resourceAppSyncFunctionConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncFunctionConfigurationType, ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncFunctionConfigurationType, ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncFunctionConfigurationType, ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncFunctionConfigurationType, data, meta)
}

func resourceAppSyncFunctionConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncFunctionConfigurationType, data, meta)
}
