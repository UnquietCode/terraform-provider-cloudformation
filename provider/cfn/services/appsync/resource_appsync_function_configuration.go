// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceAppSyncFunctionConfigurationCreate,
		Read:   resourceAppSyncFunctionConfigurationRead,
		Update: resourceAppSyncFunctionConfigurationUpdate,
		Delete: resourceAppSyncFunctionConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"function_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"function_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"data_source_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
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
				ForceNew: true,
			},
		},
	}
}

func resourceAppSyncFunctionConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::FunctionConfiguration", ResourceAppSyncFunctionConfiguration(), data, meta)
}

func resourceAppSyncFunctionConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::FunctionConfiguration", data, meta)
}