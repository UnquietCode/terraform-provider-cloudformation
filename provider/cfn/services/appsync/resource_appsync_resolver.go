// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncResolverType string = "AWS::AppSync::Resolver"

var appSyncResolverProperties map[string]string = map[string]string{
	"response_mapping_template_s3_location": "ResponseMappingTemplateS3Location",
	"type_name": "TypeName",
	"pipeline_config": "PipelineConfig",
	"data_source_name": "DataSourceName",
	"request_mapping_template": "RequestMappingTemplate",
	"response_mapping_template": "ResponseMappingTemplate",
	"kind": "Kind",
	"request_mapping_template_s3_location": "RequestMappingTemplateS3Location",
	"api_id": "ApiId",
	"field_name": "FieldName",
}

func ResourceAppSyncResolver() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppSyncResolverExists,
		Read: resourceAppSyncResolverRead,
		Create: resourceAppSyncResolverCreate,
		Update: resourceAppSyncResolverUpdate,
		Delete: resourceAppSyncResolverDelete,
		CustomizeDiff: resourceAppSyncResolverCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"pipeline_config": {
				Type: schema.TypeList,
				Elem: propertyResolverPipelineConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"data_source_name": {
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
			"kind": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"field_name": {
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

func resourceAppSyncResolverExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppSyncResolverRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncResolverType, ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncResolverType, ResourceAppSyncResolver(), data, appSyncResolverProperties, meta)
}

func resourceAppSyncResolverUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncResolverType, ResourceAppSyncResolver(), data, appSyncResolverProperties, meta)
}

func resourceAppSyncResolverDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncResolverType, data, meta)
}

func resourceAppSyncResolverCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncResolverType, data, meta)
}
