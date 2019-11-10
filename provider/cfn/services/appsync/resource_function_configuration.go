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

func ResourceFunctionConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceFunctionConfigurationCreate,
		Read:   resourceFunctionConfigurationRead,
		Update: resourceFunctionConfigurationUpdate,
		Delete: resourceFunctionConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"data_source_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_mapping_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"response_mapping_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"function_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_mapping_template_s3_location": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFunctionConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::FunctionConfiguration", data, meta)
}

func resourceFunctionConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::FunctionConfiguration", data, meta)
}

func resourceFunctionConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::FunctionConfiguration", data, meta)
}

func resourceFunctionConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::FunctionConfiguration", data, meta)
}